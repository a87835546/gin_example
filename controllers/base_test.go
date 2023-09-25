package controllers

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	//for i := 15; i < 21; i++ {
	url := fmt.Sprintf("http://aosikazy.com/index.php/vod/detail/id/308948.html")
	urls := make([]string, 0)
	//id := 17
	//title := "午夜"
	c := colly.NewCollector(
		//colly.Async(true),
		colly.MaxDepth(2),
	)
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("div.detail p", func(e *colly.HTMLElement) {
		fmt.Printf("detail -->>>%s\n", e.Text)
		if strings.Contains(e.Text, "：") {
			titles := strings.Split(e.Text, "：")
			if titles[0] == "主演" {
				log.Printf("主演-->> %s", titles[1])
			}
			if titles[0] == "类型" {
				log.Printf("类型-->> %s", titles[1])
			}

			if titles[0] == "备注" {
				log.Printf("备注-->> %s", titles[1])
			}
			if titles[0] == "上映日期" {
				year := titles[1]
				//y, err := strconv.Atoi(year)
				t, err := time.Parse("2001-01-01 15:04:05", year)
				if err == nil {
					//v.Year = y
					log.Printf("year--->>%v\n", t)

				} else {
					log.Printf("err--->>%s\n", err)
				}
			}
		}
	})

	c.OnHTML("h1.limit", func(e *colly.HTMLElement) {
		fmt.Printf("title -->>>%s\n", e.Text)

	})
	c.OnHTML("img#detail-img", func(e *colly.HTMLElement) {
		fmt.Printf("image -->>>%s\n", e.Attr("src"))

	})
	c.OnHTML("div.link input", func(e *colly.HTMLElement) {
		fmt.Printf("videoName-->>>%s\n", e.Attr("value"))
		if len(e.Attr("value")) > 1 {
			urls = append(urls, e.Attr("value"))
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

	//for i := 0; i < len(urls); i++ {
	//	url1 := fmt.Sprintf("https://bfzy.tv%s", urls[i])
	//parserOne(url1, title, id)
	//}
	//}
}

func TestB(t *testing.T) {
	tests := make(chan int, 10)
	go func() {
		for {
			if i, ok := <-tests; ok {
				log.Printf("i--->>> %d\n", i)
			}
		}
	}()
	for i := 0; i < 100; i++ {
		tests <- i
	}
}

func TestC(t *testing.T) {
	err := parserOnePron("http://aosikazy.com/index.php/vod/detail/id/308948.html", "午夜", 17)
	if err != nil {
		return
	}
}
