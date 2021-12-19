package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gocolly/colly"
)

type etymology struct {
	Word      string
	Etymology string
}

func main() {
	scrape()
}

func scrape() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 5 * time.Second,
	})

	etymList := []etymology{}

	alphabetSelector := "ul.alphabet__inner--2NEtM > li.alphabet__node--huwT8 > a[href]"
	pageSelector := "ul.ant-pagination > li.ant-pagination-item > a[href]"
	etymSelector := ".word--C9UPa.word_4pc--2SZw8"

	c.OnHTML(alphabetSelector, func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})
	c.OnHTML(pageSelector, func(e *colly.HTMLElement) {
		fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})
	c.OnHTML(etymSelector, func(e *colly.HTMLElement) {
		etym := etymology{}
		etym.Word = e.DOM.Find("a.word__name--TTbAA").First().Text()
		etym.Etymology = e.DOM.Find("section.word__defination--2q7ZH").First().Text()

		etymList = append(etymList, etym)
	})

	c.Visit("https://www.etymonline.com/")
	c.Wait()

	file, _ := json.MarshalIndent(etymList, "", "  ")
	path := filepath.Join(".", "output")
	os.MkdirAll(path, os.ModePerm)
	_ = ioutil.WriteFile("output/etymologies.json", file, 0644)
}
