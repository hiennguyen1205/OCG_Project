package controller

import (
	"net/http"
	"encoding/json"
	
	"bt/project/repository"
)

func GetAllOrdersDetails(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	listOrders := repository.GetAllOrdersDetails()
	json.NewEncoder(write).Encode(listOrders)
}