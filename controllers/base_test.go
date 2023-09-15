package controllers

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"testing"
)

func TestA(t *testing.T) {
	for i := 15; i < 21; i++ {
		url := fmt.Sprintf("https://bfzy.tv/index.php/vod/type/id/29/page/%d.htmll", i)
		urls := make([]string, 0)
		id := 8
		title := "电影"
		c := colly.NewCollector(
			//colly.Async(true),
			colly.MaxDepth(2),
		)
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		})
		c.OnHTML(".videoName", func(e *colly.HTMLElement) {
			fmt.Printf("videoName-->>>%s\n", e.Attr("href"))
			url := e.Attr("href")
			if len(url) > 0 {
				urls = append(urls, e.Attr("href"))
			}
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err.Error())
		})

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", RandomString())
			fmt.Println("Visiting", r.URL.String())
		})

		c.Visit(url)

		for i := 0; i < len(urls); i++ {
			url1 := fmt.Sprintf("https://bfzy.tv%s", urls[i])
			parserOne(url1, title, id)
		}
	}
}
