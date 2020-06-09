package controllers

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test"))
}

func SaySike(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sike"))
}
