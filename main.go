package main

import (
	"AmazonPriceTracker/controllers"
	"AmazonPriceTracker/router"
	"log"
	"net/http"
)

func main() {

	r := router.New()

	r.RegisterRoute("/", controllers.Test, "GET")
	r.RegisterRoute("/sike", controllers.SaySike, "GET")
	r.RegisterRoute("/products/all", controllers.GetAll, "POST")

	log.Fatal(http.ListenAndServe(":8000", r.Router()))
}
