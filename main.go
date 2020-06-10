package main

import (
	"AmazonPriceTracker/controllers"
	"AmazonPriceTracker/router"
	"log"
	"net/http"
)

func main() {

	r := router.New()

	r.RegisterRoute("/products", controllers.GetAll, "GET")
	r.RegisterRoute("/products/new", controllers.NewProduct, "POST")

	log.Fatal(http.ListenAndServe(":8000", r.Router()))
}
