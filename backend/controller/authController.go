package controller

import (
	"bt/project/models"
	"bt/project/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

		//request.AddCookie(cookie)
		// write.Header().Set("jwt", token)
		http.SetCookie(write, cookie)
		http.SetCookie(write, cookie1)
		write.WriteHeader(http.StatusOK)

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

func GetCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("id")
	if err != nil {
		log.Println("không lấy được cookie")
		statusCode := http.StatusUnauthorized
		http.Error(w, "Token doesnt exist", statusCode)
	}

	return c
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	var result models.User
	c := GetCookie(w, r)
	log.Println(c.Value)

	intIdUser, _ := strconv.Atoi(c.Value)
	result = repository.GetUserById(intIdUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type userPassword struct {
	Password string `json:"password"`
}

func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("id")
	if err != nil {
		log.Println("không lấy được cookie")
		statusCode := http.StatusUnauthorized
		http.Error(w, "Token doesnt exist", statusCode)
	}
	requestBody, _ := ioutil.ReadAll(r.Body)
	var userpass userPassword
	json.Unmarshal(requestBody, &userpass)
	passwordencoded, _ := bcrypt.GenerateFromPassword([]byte(userpass.Password), 14)
	intIdUser, _ := strconv.Atoi(c.Value)
	err = repository.UpdateUserPasword(string(passwordencoded), intIdUser)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("Đổi mật khẩu thành công")
	}
	w.Header().Set("Content-Type", "application/json")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var user models.User
	err1 := json.Unmarshal(requestBody, &user)
	if err1 != nil {
		log.Println(err1)
	}
	err2 := repository.UpdateUser(&user)
	if err2 != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("Failed")
		statusCode := http.StatusInternalServerError
		http.Error(w, "Can't change password", statusCode)
	} else {
		json.NewEncoder(w).Encode("Success")
	}
	w.Header().Set("Content-Type", "application/json")

}
