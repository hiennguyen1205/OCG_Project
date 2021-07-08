package repository

import (
	"bt/project/models"

	"log"
)

func GetAllOrdersDetails() (result []models.OrderDetail) {
	strQuery := "SELECT * FROM orders"
	var order models.OrderDetail
	response, err := db.Query(strQuery)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err := response.Scan(&order.Id, &order.UserId, &order.TotalPrice, &order.Payment, &order.Discount)
		if err != nil {
			log.Println(err)
		}
	}
	defer response.Close()
	return result
}
