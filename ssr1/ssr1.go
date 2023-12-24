package ssr1

import (
	"fmt"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
	"net/http"
)

type Movies struct {
	url  string
	name string
}

func SSR1() {
	// 1.请求目标网站
	response, err := http.Get("https://ssr1.scrape.center/")
	if err != nil {
		return
	}
	defer response.Body.Close()
	doc, err := libxml2.ParseHTMLReader(response.Body)
	if err != nil {
		return
	}
	defer doc.Free()
	nodeList := xpath.NodeList(doc.Find(`//*[@id="index"]/div[1]/div/div/div[@class="el-card__body"]/div[@class="el-row"]`))
	for i, node := range nodeList {
		url, _ := node.Find(".//div[1]/a/img/@src")
		movieName, _ := node.Find(".//div[2]/a/h2/text()")
		tag, _ := node.Find(`.//div[2]/div[@class="categories"]/button/span/text()`)
		date, _ := node.Find(`.//div[2]/div[3]/span/text()`)
		score, _ := node.Find(`.//div[3]/p[@class="score m-t-md m-b-n-sm"]/text()`)
		fmt.Printf("%d: %v, %v, %v, %v, %v, \n", i, url, movieName, tag, date, score)

	}
}
