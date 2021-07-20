package repository

import (
	"bt/project/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) (err error) {

	strQuery := "INSERT INTO users(username, password, email, address) VALUES (?,?,?,?)"
	result, err := db.Exec(strQuery, user.Username, user.Password, user.Email, user.Address)
	result.LastInsertId()
	return err
}

func CheckValid(u *models.User) (bool, int) {
	var user models.User
	response := db.QueryRow("SELECT * FROM users WHERE username = ?", u.Username)
	if response == nil {
		fmt.Println("Lỗi database 5**")
		return false, -1
	}
	err := response.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Address, &user.Role)
	if err != nil {
		fmt.Println("Tài khoản sai")
		return false, -1
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		fmt.Println("Sai mật khẩu")
		return false, -1
	}

	return true, user.Id

}

func GetUserById(id int) models.User {
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Address, &user.Role)
	if err != nil {
		fmt.Println("Error in GetUserById()")
	}
	return user
}

func UpdateUserPasword(pass string, id int) (err error) {

	strQuery, err := db.Prepare("UPDATE users SET password = ? WHERE id=?")
	if err != nil {
		log.Println("Cannot change password")
		fmt.Println(err)
	}
	strQuery.Exec(pass, id)
	return err
}

func UpdateUser(u *models.User) (err error) {

	strQuery, err := db.Prepare("UPDATE users SET username = ?, password = ?, email = ?, address = ?, role = ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(u.Username, u.Password, u.Email, u.Address, u.Role, u.Id)
	return err
}
