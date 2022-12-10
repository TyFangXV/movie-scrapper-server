package router

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Movie struct {
	Title    string   `json:Title`
	Img      []string `json:Img`
	Url      []string `json:Url`
	Recorder string   `json:Recorder`
	Duration []string `json:Duration`
	Type     string   `json:Type`
}

type cachedMovieList struct {
	Date      string
	movieList []Movie
}

var cachedMovie cachedMovieList

func ExploreRouter(r *gin.Context) {
	var recommendationList []string

	params := r.Query("id")

	//verify the query
	if params == "" {
		fmt.Fprint(r.Writer, "NO ID")
		return
	}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	//dyanamically add the section value to parse the dom without repetion
	selector := fmt.Sprintf("section:nth-child(%s) > div.content:nth-child(2) > div.filmlist", params)

	c.OnHTML(selector, func(h *colly.HTMLElement) {

		//parse the dom and convert it into json and append it to the list
		h.ForEach("div", func(i int, moviesEl *colly.HTMLElement) {
			m := Movie{
				Title:    moviesEl.ChildText("a.title"),
				Img:      moviesEl.ChildAttrs("a.poster > img", "src"),
				Url:      moviesEl.ChildAttrs("a.poster", "href"),
				Recorder: moviesEl.ChildText("div.quality"),
				Duration: strings.Fields(moviesEl.ChildText("div.meta")),
				Type:     moviesEl.ChildText("i.type"),
			}

			if m.Title != "" {
				jsonParsed, err := json.Marshal(m)

				if err != nil {
					fmt.Println(err)
				}
				recommendationList = append(recommendationList, string(jsonParsed))
			}

		})

		fmt.Fprint(r.Writer, recommendationList)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.OnError(func(k *colly.Response, err error) {
		fmt.Fprintf(r.Writer, "Error")
	})

	c.Visit("https://fmovies.to/home")

}
