package models

type OrderItem struct {
	Id        int `json:"id"`
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderItemInOrder struct {
	OrderId  int    `json:"id"`
	Image    string `json:"image"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
