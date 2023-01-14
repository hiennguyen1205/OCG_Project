package controller

import (
	"bt/project/repository"
	"encoding/json"
	"net/http"
)

type CategoryController struct {
	CategoryRepository *repository.CategoryRepository
}

func (cc *CategoryController) GetAllCategories(write http.ResponseWriter, request *http.Request) {
	listCategories := cc.CategoryRepository.GetAllCategories()
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(listCategories)
}
