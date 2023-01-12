package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bt/project/connect"
	"bt/project/models"
	dto "bt/project/models/dto"
	"bt/project/repository"

	"github.com/gorilla/mux"
)

func GetOrdersDetailsByUserId(write http.ResponseWriter, request *http.Request) {
	//id user lưu trên cookie
	vars := mux.Vars(request)
	strIdUser, _ := strconv.Atoi(vars["user_id"])
	// log.Println(strIdUser)
	write.Header().Set("Content-Type", "application/json")
	listOrders := repository.GetOrdersDetailsByUserId(strIdUser)
	// log.Println(listOrders)
	json.NewEncoder(write).Encode(listOrders)
}

// save order to database avtive = 0
func SaveOrderByUserNotActiveController(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var saveOrder dto.DisplayOrder
	json.Unmarshal(requestBody, &saveOrder)
	// log.Println(saveOrder)
	write.Header().Set("content-type", "application/json")
	result := repository.SaveOrderByUserNotActive(saveOrder)
	json.NewEncoder(write).Encode(result)
}

//information order
func GetInformationOrder(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdUser, _ := strconv.Atoi(vars["user_id"])
	// log.Println(strIdUser)
	info := repository.InformationOrder(strIdUser)
	json.NewEncoder(write).Encode(info)
}

// save order to database avtive = 1
func SaveOrderByUserActive(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var saveOrder dto.DisplayOrder
	json.Unmarshal(requestBody, &saveOrder)
	write.Header().Set("content-type", "application/json")
	result := repository.SaveOrderByUserActive(saveOrder)
	json.NewEncoder(write).Encode(result)
}

type EmailData struct {
	User        models.User
	ListProduct []dto.DetailProduct
}

func ChangeOrderState(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var metaData dto.DisplayOrder
	json.Unmarshal(requestBody, &metaData)
	//lay thong tin user
	user := repository.GetUserById(metaData.User.UserId)
	// Lay thong tin product voi active = 0
	listProducts := metaData.Products

	emailData := EmailData{
		User:        user,
		ListProduct: listProducts,
	}

	data, err := json.Marshal(emailData)

	if err != nil {
		fmt.Println(err)
		return
	}
	connect.PublishingAMessage(data)
	json.NewEncoder(write).Encode("Success")

}

func UpdateOrderDetails(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var orderInfo dto.DisplayOrder
	json.Unmarshal(requestBody, &orderInfo)
	write.Header().Set("content-type", "application/json")
	// log.Println(orderInfo)
	result := repository.UpdateOrderDetails(orderInfo)
	json.NewEncoder(write).Encode(result)
}

func GetAllOrrderDetailsForAdmin(write http.ResponseWriter, request *http.Request) {
	result := repository.GetAllOrrderDetailsForAdmin()
	write.Header().Set("content-type", "application/json")
	json.NewEncoder(write).Encode(result)
}

func DeleteOrderDetails(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	orderId,_ := strconv.Atoi(vars["order_id"])
	result := repository.DeleteOrderDetails(orderId)
	write.Header().Set("content-type", "application/json")
	json.NewEncoder(write).Encode(result)
}

func DetailsOrderAdmin(write http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	orderId,_ := strconv.Atoi(vars["order_id"])
	result := repository.GetDetailOrder(orderId)
	write.Header().Set("content-type", "application/json")
	json.NewEncoder(write).Encode(result)
}
