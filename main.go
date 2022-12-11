package main

import (
	"fmt"
	"log"
	"os"
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

	//port
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
	fmt.Printf("Server started at %s", port)
}
