package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bt/project/models"
	"bt/project/repository"
)

func GetAllProducts(write http.ResponseWriter, request *http.Request) {
	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))
	cursor, _ := strconv.Atoi(request.URL.Query().Get("cursor"))
	categoryId, _ := strconv.Atoi(request.URL.Query().Get("categoryId"))
	search := request.URL.Query().Get("search")
	sort := request.URL.Query().Get("sort")
	isFeature,_ := strconv.Atoi(request.URL.Query().Get("isFeature"))
	allProducts,_ := strconv.Atoi(request.URL.Query().Get("all"))
	write.Header().Set("Content-Type", "application/json")
	listProducts := repository.GetAllProducts(limit, cursor, search,sort, categoryId, isFeature, allProducts)
	json.NewEncoder(write).Encode(listProducts)
}

func GetProductById(write http.ResponseWriter, request *http.Request) {
	var result models.Product
	vars := mux.Vars(request)
	strIdProduct := vars["id_product"]
	intIdProduct, _ := strconv.Atoi(strIdProduct)
	result = repository.GetProductById(intIdProduct)
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(result)
}

func CreateProduct(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var product models.Product
	json.Unmarshal(requestBody, &product)
	result := repository.CreateProduct(product)
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(result)
}

func UpdateProduct(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var product models.Product
	json.Unmarshal(requestBody, &product)
	repository.UpdateProduct(product)
	write.Header().Set("Content-Type", "application/json")
}

func DeleteProductById(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	strIdProduct := vars["id_product"]
	intIdProduct, _ := strconv.Atoi(strIdProduct)
	result, _ := repository.DeleteProductById(intIdProduct)
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(result)
}