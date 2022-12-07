package router

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

type movie struct {
	Title    string `json:Title`
	Img      string `json:Img`
	Url      string `json:Url`
	Recorder string `json:Recorder`
	Duration string `json:Duration`
	Type     string `json:Type`
}

func ExploreRouter(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector(
		colly.AllowedDomains("https://fmovies.to/home"),
	)

	c.OnHTML("#body > div.container > section:nth-child(2) > div:nth-child(2) > div > div:nth-child(1) > h3 > a", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://fmovies.to/home")

	fmt.Fprintf(w, "owkring")

}
