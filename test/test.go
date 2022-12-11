package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type MovieType struct {
	PlayerLink string
	Title      string
	Country    string
	Release    string
	Director   string
	Production string
	Duration   string
	Genre      string
}

func main() {
	scrapper := colly.NewCollector(
		colly.AllowedDomains("fmovies.to"),
	)

	scrapper.Limit(&colly.LimitRule{
		DomainGlob:  "*fmovies.to.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	scrapper.OnHTML("#body", func(dom *colly.HTMLElement) {
		movieData := MovieType{
			Title:      dom.ChildText("div.info > h1.title"),
			Country:    dom.ChildText("div.info > div.meta > div:nth-child(1) > span:nth-child(2)"),
			Genre:      dom.ChildText("div.info > div.meta > div:nth-child(2) > span:nth-child(2)"),
			Release:    dom.ChildText("div.info > div.meta > div:nth-child(3) > span:nth-child(2)"),
			Production: dom.ChildText("div.info > div.meta > div:nth-child(5) > span:nth-child(2)"),
			Director:   dom.ChildText("div.info > div.meta > div:nth-child(4) > span:nth-child(2)"),
			Duration:   dom.ChildText("div.info > > div.meta.lg > span:nth-child(3)"),
			PlayerLink: dom.ChildText("#watch > div.play > div.container > #player > iframe"),
		}

		parsedMovieData, err := json.Marshal(movieData)
		MovieData := string(parsedMovieData)
		if err != nil {
			fmt.Println("broke")
			return
		}
		fmt.Println(MovieData)
	})

	scrapper.Visit("https://fmovies.to/movie/guillermo-del-toros-pinocchio-x159q/1-full")
}
