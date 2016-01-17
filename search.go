package xvideos4go

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	homeUrl = "http://jp.xvideos.com"
	searchUrl = homeUrl + "?k=%s"
)

func Search(query []string) []Video {
	var url string
	if (len(query) == 0) {
		url = homeUrl
	} else {
		url = fmt.Sprintf(
			searchUrl,
			strings.Join(query, "+"),
		)
	}

	doc, _ := goquery.NewDocument(url)

	var arr []Video
	doc.Find("div.mozaique").Children().Each(func(_ int, s *goquery.Selection) {
		cld := s.Find("div.thumbInside")
		title := cld.Find("p a[href]").Text()
		thumbnail, _ := cld.Find("div.thumb img[src]").Attr("src")
		url, _ := cld.Find("a[href]").Attr("href")
		duration := cld.Find("span.duration").Text()
		arr = append(
			arr,
			Video{
				Title:title,
				Duration:duration,
				ThumbNail:thumbnail,
				Url:homeUrl + url,
			},
		)
	})

	return arr
}
