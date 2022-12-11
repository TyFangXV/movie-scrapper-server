package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func MovieScreenParser(router *gin.Context) {
	url := router.Query("link")

	if url == "" {
		fmt.Fprintf(router.Writer, "NO URL")
		return
	}

	scrapper := colly.NewCollector(
		colly.AllowedDomains("fmovies.to"),
	)

	scrapper.OnHTML("#body", func(dom *colly.HTMLElement) {
		println(dom.ChildText("div.info > h1.title"))
	})

	scrapper.Visit(url)
}
