package main

import (
	promotionsController "CodingTask/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/promotions/", promotionsController.GetPromotion)
	http.ListenAndServe(":8080", nil)
}
