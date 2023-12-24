package maoyan

import (
	"fmt"
	"github.com/lestrrat-go/libxml2/types"
	"github.com/lestrrat-go/libxml2/xpath"
)

// 每个数据的 xpath
const (
	ParentContent = `//*[@id="index"]/div[1]/div/div/div[@class="el-card__body"]/div[@class="el-row"]`
	ImgUrl        = `.//div[1]/a/img/@src`
	MovieName     = `.//div[2]/a/h2/text()`
	Tags          = `.//div[2]/div[@class="categories"]/button/span/text()`
	Date          = `.//div[2]/div[3]/span/text()`
	Score         = `.//div[3]/p[@class="score m-t-md m-b-n-sm"]/text()`
)

// Result 需要拿到的数据结构
type Result struct {
	Url,
	MovieName,
	Tags,
	Date,
	Score types.XPathResult
}

// ParseResult 所有拿到的数据
type ParseResult struct {
	Data []Result
}

// Parse 自定义解析函数
func Parse(doc types.Document) {
	var result = ParseResult{}
	nodeList := xpath.NodeList(doc.Find(ParentContent))
	for _, node := range nodeList {
		url, _ := node.Find(ImgUrl)
		movieName, _ := node.Find(MovieName)
		tag, _ := node.Find(Tags)
		date, _ := node.Find(Date)
		score, _ := node.Find(Score)
		result.Data = append(result.Data, Result{
			Url:       url,
			MovieName: movieName,
			Tags:      tag,
			Date:      date,
			Score:     score,
		})
	}

	for i, data := range result.Data {
		fmt.Printf("%d: %v\n", i, data)
	}
}
