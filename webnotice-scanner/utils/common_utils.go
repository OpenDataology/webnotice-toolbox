package utils

import (
	"errors"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func Fetch(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	var content string
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return content, nil
		case tt == html.TextToken:
			content += z.Token().Data
		}
	}
}

func FetchByGoqueryForLicenseCompliance(url string) (*goquery.Document, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request err:%V", err)
		return nil, err
	}

	// 添加自定义的请求头
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)

	//resp, err := http.Get(url)
	if err != nil {
		return &goquery.Document{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("FetchByGoqueryForLicenseCompliance -- request not success,url:%s", url)
		return &goquery.Document{}, errors.New("this request did not sucess.")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return &goquery.Document{}, err
	}

	return doc, nil

}

func FetchContentByColly(url string) (string, error) {

	//url := "https://www.jtexpress.com/sc/termsOfUse"
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	})

	var content string

	c.OnHTML("p,h1,h2,h3,h4,h5,h6,div", func(element *colly.HTMLElement) {
		content += strings.TrimSpace(element.Text) + "\n"
	})

	err := c.Visit(url)

	if err != nil {
		log.Printf("FetchContentByColly is err ,url:%s: %V", url, err)
		return "", err
	}

	return content, nil
}

func ItemURLResolve(baseURL, relativeURL string) string {
	// 判断 relativeURL 是否是绝对 URL
	if strings.HasPrefix(relativeURL, "http://") ||
		strings.HasPrefix(relativeURL, "https://") {
		return relativeURL
	}

	// 否则，将其解析为绝对 URL
	return baseURL + relativeURL
}

func URLResolve(sourceURL string, protocol string) string {

	if strings.HasPrefix(sourceURL, "http://") ||
		strings.HasPrefix(sourceURL, "https://") {
		return sourceURL
	}

	if protocol != "" {
		return protocol + sourceURL
	}
	// defoult https
	return "https://" + sourceURL
}
