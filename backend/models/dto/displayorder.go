package models

type DisplayOrder struct {
	User     DetailUser      `json:"user"`
	Products []DetailProduct `json:"products"`
}

type DetailProduct struct {
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Image     string `json:"image"`
	Active    int    `json:"active"`
}

type DetailUser struct {
	UserId     int     `json:"user_id"`
	OrderId    int     `json:"order_id"`
	Discount   int     `json:"discount"`
	TotalPrice float64 `json:"total_price"`
	Payment    string  `json:"payment"`
}
