package repository

import (
	// "log"

	// "bt/project/models"
	dto "bt/project/models/dto"
	"database/sql"

	// "container/list"
	"log"
)

func GetOrdersDetailsByUserId(id int) (result dto.DisplayOrder) {
	var product dto.DetailProduct
	var user dto.DetailUser
	//lấy danh sách sản phẩm của user chưa mua (active = 0)
	strQuery := "SELECT od.user_id, od.id, od.discount, od.total_price, od.payment,oi.`active`, oi.product_id, oi.quantity, p.`name`, p.price, p.image " +
		"FROM order_items oi " +
		"JOIN order_details od ON od.id = oi.order_id " +
		"JOIN products p ON p.id = oi.product_id " +
		"WHERE od.user_id = ? AND oi.`active` = 0 "
	response, err := db.Query(strQuery, id)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err := response.Scan(&user.UserId, &user.OrderId, &user.Discount, &user.TotalPrice, &user.Payment, &product.Active, &product.ProductId, &product.Quantity, &product.Name, &product.Price, &product.Image)
		if err != nil {
			log.Println(err)
		}
		result.User = user
		result.Products = append(result.Products, product)
	}
	//nếu người đó không có danh sách sản phẩm thì danh sách sản phẩm null
	if result.Products == nil {
		strQuery := "SELECT od.user_id, od.discount, od.total_price, od.payment, oi.order_id " +
			"FROM order_items oi " +
			"JOIN order_details od ON od.id = oi.order_id " +
			"WHERE od.user_id = ?"
		response := db.QueryRow(strQuery, id)
		if err != nil {
			log.Println(err)
		}
		response.Scan(&user.UserId, &user.Discount, &user.TotalPrice, &user.Payment, &user.OrderId)
		result.User = user
	}
	defer response.Close()
	return result
}

//kiểm tra xem order của user có tồn tại không
func IsValidOrderByUserId(id int) bool {
	response := db.QueryRow(`SELECT id FROM order_details WHERE user_id = ?`, id)
	var check int
	err := response.Scan(&check)
	if err == sql.ErrNoRows {
		log.Println(err)
		return false
	}
	return true
}

//kiểm tra xem product có tồn tại trong list product của order item
func IsValidProductItemByOrderId(orderId int, productId int) bool {
	response := db.QueryRow(`SELECT product_id FROM order_items WHERE order_id = ? AND product_id = ?`, orderId, productId)
	var check int
	err := response.Scan(&check)
	if err == sql.ErrNoRows {
		log.Println(err)
		return false
	}
	return true
}

//thêm sản phẩm vào giở hàng, chưa check out thì active = 0
//sau đó lưu vào database
func SaveOrderByUserNotActive(display dto.DisplayOrder) (result string) {
	user := display.User
	listProducts := display.Products
	//không có user trong order details thì insert user vào order details
	if !IsValidOrderByUserId(user.UserId) {
		strQuery := "INSERT INTO order_details(user_id, total_price, payment, discount) VALUES (?,?,?,?)"
		_, err := db.Exec(strQuery, user.UserId, user.TotalPrice, user.Payment, user.Discount)
		if err != nil {
			log.Println(err)
			return "Save User failed"
		}
	} else {
		//có user thì update trong order detail
		strQuery, err := db.Prepare("UPDATE order_details SET total_price = ?, payment = ?, discount = ? WHERE id = ? ")
		if err != nil {
			log.Println(err)
			return "Order failed"
		}
		_, err = strQuery.Exec(user.TotalPrice, user.Payment, user.Discount, user.UserId)
		// log.Println(response.RowsAffected())
		if err != nil {
			log.Println(err)
			return "Order failed"
		}
		//update danh sách product trong order items
		for _, product := range listProducts {
			//insert vào order detail
			if IsValidProductItemByOrderId(user.OrderId, product.ProductId) {
				strQuery, err := db.Prepare("INSERT INTO order_items(order_id,product_id,quantity,active) VALUES(?,?,?,0)")
				if err != nil {
					log.Println(err.Error())
				}
				strQuery.Exec(user.OrderId, product.ProductId, product.Quantity)
			}
		}
	}
	return "Successed!!"
}

// func SaveOrderByUserNotActive(display dto.DisplayOrder) (result string) {
// 	user := display.User
// 	listProducts := display.Products
	
// 		for _, product := range listProducts {
// 			//insert vào order detail
// 			if IsValidProductItemByOrderId(user.OrderId, product.ProductId) {
// 				strQuery, err := db.Prepare("INSERT INTO order_items(order_id,product_id,quantity,active) VALUES(?,?,?,0)")
// 				if err != nil {
// 					log.Println(err.Error())
// 				}
// 				strQuery.Exec(user.OrderId, product.ProductId, product.Quantity)
// 			}
// 		}
// 	return "Successed!!"
// }

//Sau khi user checkout thành công mới thì trường active product = 1
func SaveOrderByUserActive(display dto.DisplayOrder) (result string) {
	user := display.User
	listProducts := display.Products
	//không có user trong order details thì insert user vào order details
	if !IsValidOrderByUserId(user.UserId) {
		strQuery := "INSERT INTO order_details(user_id, total_price, payment, discount) VALUES (?,?,?,?)"
		_, err := db.Exec(strQuery, user.UserId, user.TotalPrice, user.Payment, user.Discount)
		if err != nil {
			log.Println(err)
			return "Order failed"
		}
	} else {
		//có user thì update trong order detail
		strQuery, err := db.Prepare("UPDATE order_details SET total_price = ?, payment = ?, discount = ? WHERE id = ? ")
		if err != nil {
			log.Println(err)
			return "Order failed"
		}
		_, err = strQuery.Exec(user.TotalPrice, user.Payment, user.Discount, user.UserId)
		// log.Println(response.RowsAffected())
		if err != nil {
			log.Println(err)
			return "Order failed"
		}
		//update danh sách product trong order items
		for _, product := range listProducts {
			// log.Println(IsValidProductItemByOrderId(user.OrderId, product.ProductId))
			//nếu có product khớp với order detail id thì update
			if IsValidProductItemByOrderId(user.OrderId, product.ProductId) {
				strQuery, err := db.Prepare("UPDATE order_items SET quantity = ?, active = 1 WHERE order_id = ? AND product_id = ? AND active = 0")
				if err != nil {
					log.Println(err.Error())
				}
				strQuery.Exec(product.Quantity, user.OrderId, product.ProductId)
			} else {
				// không có product trong order items thì insert
				strQuery, err := db.Prepare("INSERT INTO order_items(order_id, product_id, quantity, active) VALUES (?,?,?,1) ")
				if err != nil {
					log.Println(err.Error())
				}
				_, err = strQuery.Exec(user.OrderId, product.ProductId, product.Quantity)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
	return "Order successed"
}
