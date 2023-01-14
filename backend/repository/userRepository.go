package repository

import (
	"bt/project/models"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Db *sql.DB
}

func (ur *UserRepository) CreateUser(user *models.User) (err error) {

	strQuery := "INSERT INTO users(username, password, name, phone_number, email, address) VALUES (?,?,?,?,?,?)"
	result, _ := ur.Db.Exec(strQuery, user.Username, user.Password, user.Name, user.PhoneNumber, user.Email, user.Address)
	userId, err := result.LastInsertId()
	if err != nil {
		log.Println("Tạo user lôi, ", err)
		return err
	}

	//tạo order cho user luôn
	strQuery = "INSERT INTO order_details(user_id) VALUES (?)"
	_, err = ur.Db.Exec(strQuery, userId)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (ur *UserRepository) CheckValid(u *models.User) (bool, int) {
	var user models.User
	response := ur.Db.QueryRow("SELECT id ,password FROM users WHERE username = ?", u.Username)
	if response == nil {
		fmt.Println("Lỗi database 5**")
		return false, -1
	}
	err := response.Scan(&user.Id, &user.Password)
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

func (ur *UserRepository) GetUserById(id int) models.User {
	var user models.User
	err := ur.Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.PhoneNumber, &user.Email, &user.Address, &user.Role)
	if err != nil {
		fmt.Println("Error in GetUserById")
	}
	return user
}

func (ur *UserRepository) UpdateUserPasword(pass string, id int) (err error) {

	strQuery, err := ur.Db.Prepare("UPDATE users SET password = ? WHERE id=?")
	if err != nil {
		log.Println("Cannot change password")
		fmt.Println(err)
	}
	strQuery.Exec(pass, id)
	return err
}

func (ur *UserRepository) UpdateUser(u *models.User) (err error) {
	// log.Println(u)
	strQuery, err := ur.Db.Prepare("UPDATE users SET username = ?, password = ?, email = ?, address = ?, name = ?, phone_number = ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(u.Username, u.Password, u.Email, u.Address, u.Name, u.PhoneNumber, u.Id)

	return err
}
