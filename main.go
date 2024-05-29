package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/skyakashh/UrlShortner/router.go"
)

func main() {
	fmt.Println("mongo db api")
	r := router.Router()
	fmt.Println("server is getting started ...")
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":4000", handler))
	fmt.Println("listen and serve 4000")
}
