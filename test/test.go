package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	scrapper := colly.NewCollector()

	scrapper.OnHTML("#slider > div.swiper-wrapper", func(dom *colly.HTMLElement) {
		dom.DOM.Each(func(i int, innerDiv *goquery.Selection) {
			println(innerDiv.Find("div").Html())
		})
	})

	scrapper.Visit("https://fmovies.to/home")
}
