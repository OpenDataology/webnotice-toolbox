package handlers

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/daos"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/dto"
	"log"
	"strings"

	utils "github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/utils"
	"github.com/PuerkitoBio/goquery"
)

var copyrightHandlerInstance CopyrightCompalianceHandler = CopyrightHandlerImpl{}
var licenseHandlerInstance CopyrightCompalianceHandler = LicenseHandlerImpl{}

var CopyrightComplianceHandlerList = []CopyrightCompalianceHandler{copyrightHandlerInstance, licenseHandlerInstance}

// 接口
type CopyrightCompalianceHandler interface {
	Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO)
}

// 版权合规实现类-版权归属信息
type CopyrightHandlerImpl struct {
}

// 版权合规实现类-license信息
type LicenseHandlerImpl struct {
}

func (copyrightHandler CopyrightHandlerImpl) Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO) {
	// Possible copyright keywords
	keywords := []string{"©", "Copyright", "版权所有"}
	sourceUrl := cchqd.SourceUrl
	doc, err := utils.FetchByGoqueryForLicenseCompliance(sourceUrl)
	// is can access?
	if err != nil {
		log.Printf("copyrightHandler-handle - domain cont access%V, sourceUrl:", err, sourceUrl)
		return
	}

	selecttions := doc.Find("div")
	var copyrightFlagStr string
	for i := 0; i < selecttions.Length(); i++ {
		// 获取每个元素
		s := selecttions.Eq(i)

		itemText := s.Text()
		isSuccess := false
		for _, keyword := range keywords {
			if !strings.Contains(itemText, keyword) {
				continue
			}

			isSuccess = true
			//break
		}

		if isSuccess {
			copyrightFlagStr = itemText
			//break
		}
	}

	log.Printf("copyrightHandler-handle cur copyright flag is :" + copyrightFlagStr)

	return
}

func (lincenseHandler LicenseHandlerImpl) Handle(cchqd dto.CopyrightComplianceHandlerRequestDTO) {

	log.Printf("license mes handler run url:", cchqd.SourceUrl, cchqd.AibomId)
	url := cchqd.SourceUrl
	//aibomId := cchqd.AibomId
	licenseRes := matchLicenseContent(url)

	println(licenseRes.LicenseUrl + licenseRes.LicenseContent)
	log.Printf(licenseRes.LicenseUrl + licenseRes.LicenseContent)
	// update dataset

}

func matchLicenseContent(sourceUrl string) *dto.WebPageLicenseDTO {

	// sourceUrl := "https://china.findlaw.cn"
	// licenseSuffixList := []string{"/aboutus/lawnotice.html"}
	utils.URLResolve(sourceUrl, "https")

	var licenseRes *dto.WebPageLicenseDTO

	doc, err := utils.FetchByGoqueryForLicenseCompliance(sourceUrl)
	// is can access?
	if err != nil {
		log.Printf("domain cont access, sourceUrl:" + sourceUrl)
		return nil
	}

	//todo search in the metadata

	webPageLinkList := loadUrlAndKeyWordMatchRqe(sourceUrl, doc)
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

	licenseRes = sourceUrlPlusLicenseSuffixExplore(sourceUrl)

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

		doc, err := utils.FetchByGoqueryForLicenseCompliance(licenseLink)
		if err != nil {
			log.Printf("sourceUrlPlusLicenseSuffixExplore - curLink:%s find all license suffix url  Error: %v", licenseLink, err)
			continue
		}
		// 选择整个文档的纯文本
		licenseContent := doc.Text()

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

		doc, err := utils.FetchByGoqueryForLicenseCompliance(curLinkURL)

		if err != nil {
			log.Printf("webPageLicenseUrlKeywordMatch--curLinkURL:%s   Error: %v", curLinkURL, err)
			continue
		}
		// 选择整个文档的纯文本
		fullText := doc.Text()

		log.Printf("link match index：" + string(index) + "link： " + curLinkURL + " " + licenseUrlSuffix.UrlSuffix)
		return &dto.WebPageLicenseDTO{
			LicenseContent: fullText,
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

		doc, err := utils.FetchByGoqueryForLicenseCompliance(curLinkURL)

		if err != nil {
			log.Printf("webPageLicenseKeywordMatch--curLinkURL:%s  Error: %v", curLinkURL, err)
			continue
		}
		// 选择整个文档的纯文本
		fullText := doc.Selection.Text()
		//fullText := doc.Text()

		log.Printf("webPageLicenseKeywordMatch - keyword match index：" + string(index) + "link： " + curLinkURL + " " + LicenseDescKeywordDTo.Keyword)
		return &dto.WebPageLicenseDTO{
			LicenseContent: fullText,
			LicenseUrl:     curLinkURL,
		}
	}
	return nil
}

func loadUrlAndKeyWordMatchRqe(sourceUrl string, doc *goquery.Document) []dto.WebPageLink {

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
