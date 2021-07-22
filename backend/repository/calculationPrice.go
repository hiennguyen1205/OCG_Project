package repository

import "log"

func CalcPriceOrder(id int) int64 {
	var total int64
	var subTotal int64
	var discount int
	var pricesAndQuantityToCalc PricesAndQuantityToCalc
	strQuery := "SELECT oi.quantity, p.price, p.sale FROM order_details od JOIN order_items oi ON od.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE active = 0 AND od.user_id = ?"
	response, err := db.Query(strQuery, id)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err = response.Scan(&pricesAndQuantityToCalc.quantity, &pricesAndQuantityToCalc.price, &pricesAndQuantityToCalc.sale)
		if err != nil {
			log.Println(err)
		}
		subTotal += pricesAndQuantityToCalc.price * pricesAndQuantityToCalc.quantity
		discount += int(pricesAndQuantityToCalc.price*pricesAndQuantityToCalc.quantity) * pricesAndQuantityToCalc.sale / 100

	}
	total = subTotal + subTotal/10 - int64(discount)

	return total
}
