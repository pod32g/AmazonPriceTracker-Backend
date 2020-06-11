package controllers

import (
	"AmazonPriceTracker/models"
	"AmazonPriceTracker/scrapper"
	"AmazonPriceTracker/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Product = models.New()

type productRequest struct {
	URL   string `json:"URL"`
	price int    `json:"price"`
	Name  string
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := Product.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("SERVER ERROR"))
	}

	var response = map[string][]models.ProductStruct{"products": products}

	json.NewEncoder(w).Encode(response)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct productRequest
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))

	json.Unmarshal(body, &newProduct)

	price, err := scrapper.ExtractPrice(newProduct.URL)

	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "There was an error parsing the amazon URL"})
		return
	}

	name, err := scrapper.ExtractTitle(newProduct.URL)

	if err != nil {
		fmt.Println(err.Error())
	}

	product := &models.ProductStruct{
		URL:       newProduct.URL,
		Price:     price,
		Name:      name,
		UpdatedAt: time.Time{},
	}

	if err := Product.Add(product); err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "There was an error processing the request"})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"Success": "Created Succesfully"})
}
