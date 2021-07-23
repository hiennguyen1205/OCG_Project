package models

type EmailData struct {
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Address     string        `json:"address"`
	PhoneNumber string        `json:"phone_number"`
	Products    []ProductData `json:"products"`
}

type ProductData struct {
	ProductName string `json:"productname"`
	Price       int64  `json:"price"`
	Sale        int64  `json:"sale"`
	Quantity    int64  `json:"quantity"`
}
