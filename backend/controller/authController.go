package controller

import (
	"bt/project/models"
	"bt/project/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

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
		//jwt

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    user.Username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 ngày
		})
		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			json.NewEncoder(write).Encode("Không thể đăng nhập")
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			Secure:   false,
		}

		//request.AddCookie(cookie)
		http.SetCookie(write, cookie)

	} else {
		json.NewEncoder(write).Encode("Tài khoản hoặc mật khẩu không đúng")
	}
}

func Logout(write http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(write, cookie)

	json.NewEncoder(write).Encode("Success logout")

}
