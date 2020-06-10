package controllers

import (
	"AmazonPriceTracker/models"
	"AmazonPriceTracker/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var Product = models.New()

type productRequest struct {
	URL   string
	price string
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := Product.GetAll()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("SERVER ERROR"))
	}

	json.NewEncoder(w).Encode(products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {

	var productRqst productRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&productRqst); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request", "message": err.Error()})
		return
	}
	fmt.Println(productRqst)
	defer r.Body.Close()

	price, err := strconv.ParseFloat(productRqst.price, 64)

	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request", "message": err.Error()})
		return
	}

	product := &models.ProductStruct{
		URL:       productRqst.URL,
		Price:     price,
		UpdatedAt: time.Now(),
	}

	if err := Product.Add(product); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "There was an error processing the request"})
		return
	}

}
