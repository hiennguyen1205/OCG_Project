package repository

import (
	"bt/project/models"
	dto "bt/project/models/dto"
	"database/sql"

	"log"
)

type OrderRepository struct {
	Db *sql.DB
}

// thông tin đơn hàng
func (or *OrderRepository) InformationOrder(id int) (result dto.DisplayOrder) {
	var product dto.DetailProduct
	var user dto.DetailUser
	//tính tiền
	//lấy danh sách sản phẩm của user đã mua (active = 1)
	strQuery := "SELECT od.user_id, od.id, od.discount, od.total_price, oi.`active`, oi.product_id, oi.quantity, p.`name`, p.price " +
		"FROM order_items oi " +
		"JOIN order_details od ON od.id = oi.order_id " +
		"JOIN products p ON p.id = oi.product_id " +
		"WHERE od.user_id = ? AND oi.`active` = 1"
	response, err := or.Db.Query(strQuery, id)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err := response.Scan(&user.UserId, &user.OrderId, &user.Discount, &user.TotalPrice, &product.Active, &product.ProductId, &product.Quantity, &product.Name, &product.Price)
		if err != nil {
			log.Println(err)
		}
		result.User = user
		result.Products = append(result.Products, product)
	}
	defer response.Close()
	return result
}

func (or *OrderRepository) GetOrdersDetailsByUserId(id int) (result dto.DisplayOrder) {
	var product dto.DetailProduct
	var user dto.DetailUser
	price := 0
	payment := ""
	//lấy danh sách sản phẩm của user chưa mua (active = 0)
	strQuery := "SELECT od.user_id, od.id, od.discount, oi.`active`, oi.product_id, oi.quantity, p.`name`, p.price, p.image, p.sale " +
		"FROM order_items oi " +
		"JOIN order_details od ON od.id = oi.order_id " +
		"JOIN products p ON p.id = oi.product_id " +
		"WHERE od.user_id = ? AND oi.`active` = 0 AND od.is_delete = 0"
	response, err := or.Db.Query(strQuery, id)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err := response.Scan(&user.UserId, &user.OrderId, &user.Discount, &product.Active, &product.ProductId, &product.Quantity, &product.Name, &product.Price, &product.Image, &product.Sale)
		if err != nil {
			log.Println(err)
		}
		result.User = user
		result.Products = append(result.Products, product)
	}

	//nếu người đó không có danh sách sản phẩm thì danh sách sản phẩm null
	if result.Products == nil {
		strQuery := "SELECT od.user_id, oi.order_id " +
			"FROM order_items oi " +
			"JOIN order_details od ON od.id = oi.order_id " +
			"WHERE od.user_id = ?"
		response := or.Db.QueryRow(strQuery, id)
		err := response.Scan(&user.UserId, &user.OrderId)
		if err == sql.ErrNoRows {
			log.Println(err)
			strQuery = "SELECT * FROM order_details WHERE user_id = ? "
			response = or.Db.QueryRow(strQuery, id)
			response.Scan(&user.OrderId, &user.UserId, &price, &payment, &user.Discount)
		}
		result.User = user
	}
	defer response.Close()
	return result
}

// kiểm tra xem order của user có tồn tại không
func (or *OrderRepository) IsValidOrderByUserId(id int) bool {
	response := or.Db.QueryRow(`SELECT id FROM order_details WHERE user_id = ?`, id)
	var check int
	err := response.Scan(&check)
	if err == sql.ErrNoRows {
		log.Println("không có order của user, ", err)
		return false
	}
	return true
}

// kiểm tra xem product có tồn tại trong list product của order item
func (or *OrderRepository) IsValidProductItemByOrderId(orderId int, productId int) bool {
	response := or.Db.QueryRow(`SELECT product_id FROM order_items WHERE order_id = ? AND product_id = ?`, orderId, productId)
	var check int
	err := response.Scan(&check)
	if err == sql.ErrNoRows {
		log.Println("không có product trong order,", err, " sẽ cập nhật luôn!!")
		return false
	}
	return true
}

// thêm sản phẩm vào giở hàng, chưa check out thì active = 0 sau đó lưu vào database
func (or *OrderRepository) SaveOrderByUserNotActive(display dto.DisplayOrder) (result string) {
	var (
		discount int64
		price    int64
	)
	user := display.User
	listProducts := display.Products
	//update danh sách product trong order items
	for _, product := range listProducts {
		//insert vào order items
		if !or.IsValidProductItemByOrderId(user.OrderId, product.ProductId) {
			strQuery, err := or.Db.Prepare("INSERT INTO order_items(order_id,product_id,quantity,`active`) VALUES(?,?,?,0)")
			if err != nil {
				log.Println(err.Error())
			}
			_, err = strQuery.Exec(user.OrderId, product.ProductId, product.Quantity)
			if err != nil {
				log.Println("insert vào db lỗi", err.Error())
			}
		} else {
			response := or.Db.QueryRow("SELECT active FROM order_items WHERE order_id = ? AND product_id = ? ORDER BY id DESC LIMIT 1", user.OrderId, product.ProductId)
			var check int
			err := response.Scan(&check)
			if err == sql.ErrNoRows {
				log.Println(err)
			}
			//nếu active = 1 thì insert sản phẩm với active = 0
			if check == 1 {
				strQuery, err := or.Db.Prepare("INSERT INTO order_items(order_id,product_id,quantity,active) VALUES(?,?,?,0)")
				if err != nil {
					log.Println(err.Error())
				}
				strQuery.Exec(user.OrderId, product.ProductId, product.Quantity)
			} else { //nếu active = 0 thì update số lượng
				strQuery, err := or.Db.Prepare("UPDATE order_items SET quantity = ? WHERE order_id = ? AND product_id = ? AND active = 0")
				if err != nil {
					log.Println(err.Error())
				}
				_, err = strQuery.Exec(product.Quantity, user.OrderId, product.ProductId)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}

	//update tổng tiền và tổng só tiền giảm giá
	price, discount = or.CalcAmount(user.UserId)
	_, err := or.Db.Exec("UPDATE order_details SET total_price = ?, discount = ? WHERE id = ?", price, discount, user.UserId)
	if err != nil {
		log.Println(err)
		return "Order failed"
	}
	return "Successed!!"
}

// Sau khi user checkout thành công mới thì trường active product = 1
func (or *OrderRepository) SaveOrderByUserActive(display dto.DisplayOrder) (result string) {
	user := display.User
	listProducts := display.Products
	//update danh sách product trong order items active = 1
	for _, product := range listProducts {
		strQuery, err := or.Db.Prepare("UPDATE order_items SET quantity = ?, active = 1 WHERE order_id = ? AND product_id = ?")
		if err != nil {
			log.Println(err.Error())
		}
		strQuery.Exec(product.Quantity, user.OrderId, product.ProductId)
	}
	return "Order successed"
}

func (or *OrderRepository) UpdateOrderDetails(display dto.DisplayOrder) (result string) {
	user := display.User
	log.Println("user nè ", user)

	_, err := or.Db.Exec("UPDATE order_details SET is_paied = 1 WHERE id = ? AND user_id = ?", user.OrderId, user.UserId)
	if err != nil {
		log.Println(err)
		return
	}
	return "Update Order Success!"
}

func (or *OrderRepository) GetAllOrrderDetailsForAdmin() (result []models.AdminOrderDetails) {
	rows, _ := or.Db.Query("SELECT od.id, od.total_price, od.is_paied, od.status, u.name FROM order_details od JOIN users u ON u.id = od.user_id WHERE is_delete = 0 LIMIT 20 OFFSET 0")
	var OrderDetails models.AdminOrderDetails
	for rows.Next() {
		err := rows.Scan(&OrderDetails.OrderId, &OrderDetails.TotalPrice, &OrderDetails.Ispaied, &OrderDetails.Status, &OrderDetails.UserName)
		if err != nil {
			log.Println("Lỗi không lấy được thông tin order details")
		}
		result = append(result, OrderDetails)
	}
	return result
}

func (or *OrderRepository) DeleteOrderDetails(orderId int) string {
	_, err := or.Db.Exec("UPDATE order_details SET is_delete = 1 WHERE id = ?", orderId)
	if err != nil {
		log.Println("Không xóa được order!!")
		return err.Error()
	}
	return "Xóa Order thành công"
}

func (or *OrderRepository) GetDetailOrder(orderId int) (result []models.OrderItemInOrder) {
	rows, _ := or.Db.Query("SELECT od.id, p.image, p.name, p.price, oi.quantity FROM order_items oi JOIN order_details od ON od.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE od.id = ? ", orderId)
	var items models.OrderItemInOrder
	for rows.Next() {
		err := rows.Scan(&items.OrderId, &items.Image, &items.Name, &items.Price, &items.Quantity)
		if err != nil {
			log.Println("Không hiển thị được order")
		}
		result = append(result, items)
	}
	return result
}

type PricesAndQuantityToCalc struct {
	price    int64
	quantity int64
	sale     int
}

func (or *OrderRepository) CalcAmount(id int) (total int64, discount int64) {
	var subTotal int64
	total = 0
	discount = 0
	var pricesAndQuantityToCalc PricesAndQuantityToCalc
	strQuery := "SELECT oi.quantity, p.price, p.sale FROM order_details od JOIN order_items oi ON od.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE active = 0 AND od.user_id = ?"
	response, err := or.Db.Query(strQuery, id)
	if err != nil {
		log.Println(err)
	}
	for response.Next() {
		err = response.Scan(&pricesAndQuantityToCalc.quantity, &pricesAndQuantityToCalc.price, &pricesAndQuantityToCalc.sale)
		if err != nil {
			log.Println(err)
		}
		subTotal += pricesAndQuantityToCalc.price * pricesAndQuantityToCalc.quantity
		discount += (pricesAndQuantityToCalc.price * pricesAndQuantityToCalc.quantity * int64(pricesAndQuantityToCalc.sale)) / 100
	}
	total = subTotal + subTotal/10 - int64(discount)

	return total, discount
}
