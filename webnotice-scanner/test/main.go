package main

import (
	"fmt"
	handlers "github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/handlers"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/dto"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/utils"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	// 输出字符串
	fmt.Println("Hello, World!")
	//crawingByCollyTest()
	//duotaitest()
	//collyTest()

	fetchTest()

}

func crawingByCollyTest() {
	// 创建一个新的 Colly 收集器
	c := colly.NewCollector()

	// 存储所有文本的切片
	var allText []string

	// 在收到 HTML 数据时调用回调函数
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// 提取文本并将其添加到切片中
		text := strings.TrimSpace(e.Text)
		if text != "" {
			allText = append(allText, text)
		}
	})

	// 在访问链接时调用回调函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// 启动爬虫，访问指定网址
	err := c.Visit("https://www.jtexpress.com/sc/termsOfUse")
	if err != nil {
		log.Fatal(err)
	}

	// 打印所有提取到的文本
	for _, text := range allText {
		fmt.Println(text)
	}
}

func duotaitest() {

	copyrightCompalianceHandlerQueset := dto.CopyrightComplianceHandlerRequestDTO{
		AibomId:   1,
		SourceUrl: "https://china.findlaw.cn",
	}

	//遍历
	copyrightCompalianceHandlerList := handlers.CopyrightComplianceHandlerList
	for i := 0; i < len(copyrightCompalianceHandlerList); i++ {
		copyrightCompalianceHandlerList[i].Handle(copyrightCompalianceHandlerQueset)
	}
}

func collyTest() {

	url := "https://www.jtexpress.com/sc/termsOfUse"
	c := colly.NewCollector()

	var content string

	c.OnHTML("p,h1,h2,h3,h4,h5,h6,div", func(element *colly.HTMLElement) {
		content += strings.TrimSpace(element.Text) + "\n"
	})

	err := c.Visit(url)

	if err != nil {
		return
	}

	println(content)

	return
}

func fetchTest() {

	doc, _ := utils.FetchByGoqueryForLicenseCompliance("https://www.jtexpress.com/sc/termsOfUse")

	finds := doc.Find("p,h1,h2,h3,h4,h5,h6,div")

	println(finds.Text())
}
