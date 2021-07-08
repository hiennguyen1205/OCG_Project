package controller

import (
	"bt/project/models"
	"bt/project/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	// if err1 != nil {
	// 	json.NewEncoder(write).Encode(err1)
	// }
	var user models.User
	json.Unmarshal(requestBody, &user)
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(password)
	err := repository.CreateUser(&user)
	if err != nil {
		json.NewEncoder(write).Encode(err)
	} else {
		json.NewEncoder(write).Encode("Thêm thành công")
	}

}

func Login(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var user models.User
	json.Unmarshal(requestBody, &user)

	isValid := repository.CheckValid(&user)
	if isValid {
		fmt.Println("đăng nhập thành công")
	} else {
		fmt.Println("Tài khoản hoặc mật khẩu không đúng")
	}

}
