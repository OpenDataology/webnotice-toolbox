package handlers

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/daos"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/do"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/dto"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/enums"
	"log"
	"regexp"
	"sort"
	"strings"

	utils "github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/utils"
	"github.com/PuerkitoBio/goquery"
)

var copyrightHandlerInstance CopyrightComplianceHandler = CopyrightHandlerImpl{}
var licenseHandlerInstance CopyrightComplianceHandler = LicenseHandlerImpl{}

var CopyrightComplianceHandlerList = []CopyrightComplianceHandler{copyrightHandlerInstance, licenseHandlerInstance}

// interface
type CopyrightComplianceHandler interface {
	Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO)
}

// CopyrightComplianceHandler Impl-Copyright belong to
type CopyrightHandlerImpl struct {
}

// CopyrightComplianceHandler Impl-license
type LicenseHandlerImpl struct {
}

type CopyrightFlagLoader interface {
	Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO)
}

func (copyrightHandler CopyrightHandlerImpl) Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO) {
	// Possible copyright keywords
	copyrightFlag := matchCopyrightFlag(cchqd.SourceUrl)

	if copyrightFlag == nil {
		log.Printf("copyrightHandler-handle can not load copyright flag form url:" + cchqd.SourceUrl)
		return
	}

	log.Printf("copyrightHandler-handle cur copyright flag is :" + copyrightFlag.Flag)

	err := daos.LicenseMesSave(do.LicenseMes{
		LicenseName:   cchqd.SourceUrl,
		CopyrightFlag: copyrightFlag.Flag,
		Licensor:      copyrightFlag.Licensor,
		SourceUrl:     cchqd.SourceUrl,
		AibomId:       cchqd.AibomId,
	})

	if err != nil {
		log.Printf("copyrightHandler-handle: save license data err:%V", err)
	}

	return
}

func (licenseHandler LicenseHandlerImpl) Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO) {

	log.Printf("license mes handler run url:%s,bimId:%d", cchqd.SourceUrl, cchqd.AibomId)
	//aibomId := cchqd.AibomId
	licenseRes := matchLicenseContent(cchqd.SourceUrl)

	log.Printf("LicenseHandlerImpl-handler result url:%s,content:%s",
		licenseRes.LicenseUrl, licenseRes.LicenseContent)
	// update data
	err := daos.LicenseMesSave(do.LicenseMes{
		LicenseName:    cchqd.SourceUrl,
		LicenseUrl:     licenseRes.LicenseUrl,
		LicenseType:    enums.DataSourceTermsOfUse,
		LicenseContent: licenseRes.LicenseContent,
		SourceUrl:      cchqd.SourceUrl,
		AibomId:        cchqd.AibomId,
	})

	if err != nil {
		log.Printf("LicenseHandlerImpl-handler: save license data err:%V", err)
	}

}

func matchCopyrightFlag(url string) *dto.WebPageCopyrightFlagDTO {

	//todo match in the database

	//keywords := []string{"©", "Copyright", "版权所有"}
	sourceWebPageCopyrightKeywordList, err := daos.SourceWebPageCopyrightKeywordFindAll()

	sourceUrl := url
	doc, err := utils.FetchByGoqueryForLicenseCompliance(sourceUrl)
	// is can access?
	if err != nil {
		log.Printf("copyrightHandler-handle - domain cont access%V, sourceUrl:%s", err, sourceUrl)
		return nil
	}

	selections := doc.Find("div,li")
	copyrightFlagStrList := []string{}
	for i := 0; i < selections.Length(); i++ {
		// 获取每个元素
		s := selections.Eq(i)
		itemText := s.Text()

		for _, sourceWebPageCopyrightKeyword := range sourceWebPageCopyrightKeywordList {
			if !strings.Contains(itemText, sourceWebPageCopyrightKeyword.Keyword) {
				continue
			}
			copyrightFlagStrList = append(copyrightFlagStrList, itemText)
		}

	}

	if len(copyrightFlagStrList) <= 0 {
		return nil
	}
	// str length sort.Slice
	sort.Slice(copyrightFlagStrList, func(i, j int) bool {
		return len(copyrightFlagStrList[i]) < len(copyrightFlagStrList[j])
	})
	copyrightFlagStr := copyrightFlagStrList[0]

	licensorRes := extractLicensorFormCopyrightFlag(copyrightFlagStr, sourceWebPageCopyrightKeywordList)

	return &dto.WebPageCopyrightFlagDTO{
		Flag:     copyrightFlagStr,
		Licensor: licensorRes,
	}

}

func extractLicensorFormCopyrightFlag(copyrightFlagStr string, sourceWebPageCopyrightKeywordList []do.SourceWebPageCopyrightKeyword) string {

	licensorRes := copyrightFlagStr
	// remove copyrightKeyword
	for _, copyrightKeyword := range sourceWebPageCopyrightKeywordList {
		//licensorRes = strings.NewReplacer(licensorRes, "").Replace(copyrightKeyword.Keyword)
		licensorRes = strings.ReplaceAll(licensorRes, copyrightKeyword.Keyword, "")
	}

	// remove time
	regexpStrList := []string{`\b\d{4}-\d{4}\b`, `\b\d{4}\b`}
	for _, regexpStr := range regexpStrList {
		re := regexp.MustCompile(regexpStr)
		licensorRes = re.ReplaceAllString(licensorRes, "")
	}

	return licensorRes
}

func matchLicenseContent(sourceUrl string) *dto.WebPageLicenseDTO {

	// sourceUrl := "https://china.findlaw.cn"
	// licenseSuffixList := []string{"/aboutus/lawnotice.html"}
	utils.URLResolve(sourceUrl, "https")

	var licenseRes *dto.WebPageLicenseDTO = nil

	doc, err := utils.FetchByGoqueryForLicenseCompliance(sourceUrl)
	// is can access?
	if err != nil {
		log.Printf("domain cont access, sourceUrl:" + sourceUrl)
		return nil
	}

	//todo search in the metadata

	webPageLinkList := loadUrlAndKeyWordFromWebPage(sourceUrl, doc)
	for _, value := range webPageLinkList {

		// text match
		licenseRes = webPageLicenseKeywordMatch(value)

		if licenseRes != nil {
			break
		}

		// license url match
		licenseRes = webPageLicenseUrlKeywordMatch(value)

		if licenseRes != nil {
			break
		}
	}

	if licenseRes == nil {
		//try sourceUrl + license url suffix
		licenseRes = sourceUrlPlusLicenseSuffixExplore(sourceUrl)
	}

	return licenseRes

}

func sourceUrlPlusLicenseSuffixExplore(sourceUrl string) *dto.WebPageLicenseDTO {
	//licenseSuffixList := []string{"/aboutus/lawnotice.html"}
	licenseUrlSuffixList, err := daos.LicenseUrlSuffixFindAll()

	if err != nil {
		log.Printf("find all license suffix url  Error: %v", err)
		return nil
	}
	// 使用 range 遍历数组
	for _, licenseSuffix := range licenseUrlSuffixList {

		licenseLink := sourceUrl + licenseSuffix.UrlSuffix

		licenseContent, err := utils.FetchContentByColly(licenseLink)

		//doc, err := utils.FetchByGoqueryForLicenseCompliance(licenseLink)
		if err != nil {
			log.Printf("sourceUrlPlusLicenseSuffixExplore - curLink:%s find all license suffix url  Error: %v", licenseLink, err)
			continue
		}
		//// 选择整个文档的纯文本
		//licenseContent := doc.Text()

		if licenseContent != "" {
			return &dto.WebPageLicenseDTO{
				LicenseContent: licenseContent,
				LicenseUrl:     licenseLink,
			}
		}
	}

	return nil

}

func webPageLicenseUrlKeywordMatch(webPageLink dto.WebPageLink) *dto.WebPageLicenseDTO {
	// curLinkText := webPageLink.Text
	curLinkURL := webPageLink.URL

	//LicenseUrlKeywords := []string{"/aboutus/lawnotice.html", "/terms", "/chn/copyright.html", "/terms.html", "/service.html", "/fe/protocol"}
	licenseUrlSuffixList, err := daos.LicenseUrlSuffixFindAll()

	if err != nil {
		log.Printf("find all license suffix url  Error: %v", err)
		return nil
	}

	for index, licenseUrlSuffix := range licenseUrlSuffixList {
		if !strings.Contains(curLinkURL, licenseUrlSuffix.UrlSuffix) {
			continue
		}

		licenseContent, err := utils.FetchContentByColly(curLinkURL)

		if err != nil {
			log.Printf("webPageLicenseUrlKeywordMatch--curLinkURL:%s   Error: %v", curLinkURL, err)
			continue
		}
		// 选择整个文档的纯文本

		log.Printf("link match index：%d link：%s , suffix:%s ", index, curLinkURL, licenseUrlSuffix.UrlSuffix)
		return &dto.WebPageLicenseDTO{
			LicenseContent: licenseContent,
			LicenseUrl:     curLinkURL,
		}
	}
	return nil

}

func webPageLicenseKeywordMatch(webPageLink dto.WebPageLink) *dto.WebPageLicenseDTO {

	curLinkText := webPageLink.Text
	curLinkURL := webPageLink.URL

	//LicenseDescKeywords := []string{"法律声明", "服务协议", "使用条款", "使用协议", "服务条款", "版权说明", "网站声明", "terms of use"}
	LicenseDescKeywords, err := daos.LicenseKeywordFindAll()
	if err != nil {
		log.Printf("webPageLicenseKeywordMatch - find all license keyword  Error: %v", err)
		return nil
	}

	for index, LicenseDescKeywordDTo := range LicenseDescKeywords {
		if !strings.Contains(curLinkText, LicenseDescKeywordDTo.Keyword) {
			continue
		}

		licenseContent, err := utils.FetchContentByColly(curLinkURL)

		if err != nil {
			log.Printf("webPageLicenseKeywordMatch--curLinkURL:%s  Error: %v", curLinkURL, err)
			continue
		}
		// 选择整个文档的纯文本

		log.Printf("webPageLicenseKeywordMatch - keyword match index：%d,curLinkUrl:%s,keyword:%s", index, curLinkURL, LicenseDescKeywordDTo.Keyword)
		return &dto.WebPageLicenseDTO{
			LicenseContent: licenseContent,
			LicenseUrl:     curLinkURL,
		}
	}
	return nil
}

func loadUrlAndKeyWordFromWebPage(sourceUrl string, doc *goquery.Document) []dto.WebPageLink {

	var webPageLinks []dto.WebPageLink

	selectionItemA := doc.Find("a")
	for i := 0; i < selectionItemA.Length(); i++ {
		// 获取每个元素
		s := selectionItemA.Eq(i)

		linkText := s.Text()
		link, exists := s.Attr("href")

		if !exists {
			continue
		}

		link = utils.ItemURLResolve(sourceUrl, link)
		webPageLink := dto.WebPageLink{
			Text: linkText,
			URL:  link,
		}
		webPageLinks = append(webPageLinks, webPageLink)
	}

	return webPageLinks
}
