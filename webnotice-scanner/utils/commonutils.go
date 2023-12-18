package utils

import (
	"errors"
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
	resp, err := http.Get(url)
	if err != nil {
		return &goquery.Document{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("FetchByGoqueryForLicenseCompliance -- request not success,url:", url)
		return &goquery.Document{}, errors.New("this request did not sucess.")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return &goquery.Document{}, err
	}

	return doc, nil

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
