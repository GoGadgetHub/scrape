package ssr1

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/types"
	"io"
	"log"
	"net/http"
)

// Fetch 发起请求
func (r *Req) Fetch() types.Document {
	client := &http.Client{}
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get error is ", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Http status code is ", res.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	doc, err := libxml2.ParseHTMLReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
