package controllers

import (
	"fmt"
	"net/http"
)

func Another(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hitted: Another")

	w.Write([]byte("Another"))
}
