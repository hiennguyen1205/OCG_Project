package controller

import (
	"bt/project/models"
	"bt/project/repository"
	"encoding/json"
	"io/ioutil"

	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(write http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
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

	isValid, userId := repository.CheckValid(&user)
	if isValid {
		//jwt
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    user.Username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 ngày
		})
		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			statusCode := http.StatusUnauthorized
			http.Error(write, strconv.Itoa(userId), statusCode)
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Path:     "/",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		}

		cookie1 := &http.Cookie{
			Name:    "id",
			Path:    "/",
			Value:   strconv.Itoa(userId),
			Expires: time.Now().Add(time.Hour * 24),
			//HttpOnly: true,
		}

		http.SetCookie(write, cookie)
		http.SetCookie(write, cookie1)

		json.NewEncoder(write).Encode(userId)

	} else {
		statusCode := http.StatusUnauthorized
		http.Error(write, strconv.Itoa(userId), statusCode)
	}
}

func Logout(write http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	cookie1 := &http.Cookie{
		Name:    "id",
		Path:    "/",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		//HttpOnly: true,
	}

	http.SetCookie(write, cookie)
	http.SetCookie(write, cookie1)

	json.NewEncoder(write).Encode("Success logout")

}