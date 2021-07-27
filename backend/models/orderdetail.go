package models

type OrderDetail struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Payment    string  `json:"payment"`
	Discount   int     `json:"discount"`
	Active     int     `json:"active"`
	Status     string  `json:"status"`
	SendEmail  int     `json:"send_email"`
	Ispaied    int     `json:"is_paied"`
}

type AdminOrderDetails struct {
	OrderId    int    `json:"order_id"`
	TotalPrice int    `json:"total_price"`
	Ispaied    int    `json:"is_paied"`
	Status     string `json:"status"`
	UserName   string `json:"username"`
}
