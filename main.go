package main

import (
	"AmazonPriceTracker/controllers"
	"AmazonPriceTracker/router"
	"log"
	"net/http"
)

func main() {

	r := router.New()

	r.RegisterRoute("/", controllers.Test)
	r.RegisterRoute("/sike", controllers.SaySike)

	log.Fatal(http.ListenAndServe(":8000", r.Router()))
}
