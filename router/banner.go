package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type banner struct {
	Title    string   `json:Title`
	Img      []string `json:Img`
	Url      []string `json:Url`
	Recorder string   `json:Recorder`
	Duration []string `json:Duration`
	Type     string   `json:Type`
}

func BannerElement(gin *gin.Context) {
	scrapper := colly.NewCollector()

	scrapper.OnHTML("#slider > div.swiper-wrapper", func(dom *colly.HTMLElement) {
		println(dom.DOM)
	})

	scrapper.Visit("https://fmovies.to")
}
