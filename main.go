package main

import (
	"fmt"
	"log"
	"net/http"
	"yt-dl-server/router"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "200")
	})

	http.HandleFunc("/explore", router.ExploreRouter)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
