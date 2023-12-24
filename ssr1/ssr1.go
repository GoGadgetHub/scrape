package ssr1

import "scrape/ssr1/maoyan"

func SSR1() {
	var r = Req{
		Url:    "https://ssr1.scrape.center/",
		Method: "GET",
	}
	doc := r.Fetch()
	defer doc.Free()
	maoyan.Parse(doc)
}
