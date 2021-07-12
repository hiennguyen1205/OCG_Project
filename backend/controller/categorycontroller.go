package controller

import (
	"encoding/json"
	"net/http"
	
	"bt/project/repository"
)


func GetAllCategories(write http.ResponseWriter, request *http.Request) {
	listCategories := repository.GetAllCategories()
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(listCategories)
}