package models

type OrderDetail struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Payment    string  `json:"payment"`
	Discount   int     `json:"discount"`
}