package controller

import (
	utils "CodingTask/utils"
	"encoding/json"
	"net/http"
)

func GetPromotion(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := utils.ParseIdFromUrl(request.URL.Path)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(err.Error()))
		return
	}

	promo, err := utils.ParseCsvLine(id)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte(err.Error()))
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(promo)
}
