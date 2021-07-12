package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dto "bt/project/models/dto"
	"bt/project/repository"
)

func GetOrdersDetailsByUserId(write http.ResponseWriter, request *http.Request) {
	//id user luwu tren header vd id =1
	write.Header().Set("Content-Type", "application/json")
	listOrders := repository.GetOrdersDetailsByUserId(1)
	json.NewEncoder(write).Encode(listOrders)
}

//check valid order
func CheckValid(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	listOrders := repository.IsValidOrderByUserId(3)
	json.NewEncoder(write).Encode(listOrders)
}

// save order to database
func SaveOrderByUserActive(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var saveOrder dto.DisplayOrder
	json.Unmarshal(requestBody, &saveOrder)
	write.Header().Set("content-type", "application/json")
	result := repository.SaveOrderByUserActive(saveOrder)
	json.NewEncoder(write).Encode(result)
}
