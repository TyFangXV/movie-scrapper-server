package main

import (
	"fmt"
	"log"
	"yt-dl-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Println(c.Query("ID"))
		fmt.Fprintf(c.Writer, "200")
	})

	r.GET("/explore", router.ExploreRouter)

	if err := r.Run(":8080"); err != nil {
		log.Panicf("error: %s", err)
	}
	fmt.Println("Server started at 8080")
}
