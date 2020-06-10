package controllers

import (
	"AmazonPriceTracker/models"
	"encoding/json"
	"net/http"
)

var Product = models.New()

func GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := Product.GetAll()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("SERVER ERROR"))
	}

	json.NewEncoder(w).Encode(products)
}
