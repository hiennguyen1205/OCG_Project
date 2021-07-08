package repository

import (
	"bt/project/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) (err error) {

	strQuery := "INSERT INTO users(username, password, email) VALUES (?,?,?)"
	result, err := db.Exec(strQuery, user.Username, user.Password, user.Email)
	result.LastInsertId()
	return err
}

func CheckValid(u *models.User) bool {
	var user models.User
	response := db.QueryRow("SELECT * FROM users WHERE username = ?", u.Username)
	if response == nil {
		fmt.Println("Lỗi database 5**")
		return false
	}
	err := response.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)
	if err != nil {
		fmt.Println("Tài khoản sai")
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		fmt.Println("Sai mật khẩu")
		return false
	}

	return true

}
