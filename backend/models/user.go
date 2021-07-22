package models

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Role        int    `json:"role"`
}
