package repository

import (
	"fmt"
	"database/sql"

	"bt/project/connect"
	"bt/project/models"
)

// biến dùng để connect tới database
var db = connect.Connect()

//hàm insert data và database
func InitData() {
	defer db.Close()
	listCategories := []models.Category{
		{Name: "Loại 1"},
		{Name: "Loại 2"},
		{Name: "Loại 3"},
	}
	listProducts := []models.Product{
		{
			Name:        "A",
			Description: "Description A",
			Price:       340,
			CategoryId:  1,
			Image:       "Hahaha",
			Sale:        30,
			IsFeature:   false,
		},
		{
			Name:        "BA",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   true,
		},
		{
			Name:        "CB",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   false,
		},
		{
			Name:        "DBA",
			Description: "Description B",
			Price:       40,
			CategoryId:  3,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   true,
		},
		{
			Name:        "EFG",
			Description: "Description B",
			Price:       40,
			CategoryId:  3,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   true,
		},
		{
			Name:        "AVF",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   false,
		},
		{
			Name:        "GCE",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   true,
		},
		{
			Name:        "HBD",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
			IsFeature:   true,
		},
	}

	for _, category := range listCategories {
		strQuery := "INSERT INTO categories(name) VALUES (?)"
		result, err := db.Exec(strQuery, category.Name)
		if err != nil {
			panic(err.Error())
		}
		lastCategoryID, _ := result.LastInsertId()
		fmt.Printf("Danh mục mới thêm có ID là: %d\n", lastCategoryID)
	}

	for _, product := range listProducts {
		strQuery := "INSERT INTO products(name, description, price, category_id, image, sale, is_feature) VALUES (?,?,?,?,?,?,?)"
		result, err := db.Exec(strQuery, product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature)
		if err != nil {
			panic(err.Error())
		}
		lastProductId, _ := result.LastInsertId()
		fmt.Printf("Sản phẩm mới thêm có ID là: %d\n", lastProductId)
	}
}

/*
* Hàm lấy danh sách sản phẩm
* input: 4 params giúp cho việc phân trang, search theo tên và sắp xếp sản phẩm
* output: 1 slice các sản phẩm
*/
func GetAllProducts(limit int, cursor int, search string, sort string, categoryId int, isFeature int) (result []models.Product) {
	var strQuery string
	var response *sql.Rows
	lenSearch := len(search)
	lenSort := len(sort)
	if lenSearch != 0 && lenSort != 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? AND name LIKE ? ORDER BY price " + sort + " LIMIT ?")
		response, _ = db.Query(strQuery, cursor, categoryId, "%"+search+"%", limit)
	} else if lenSearch != 0 && len(sort) == 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? AND name LIKE ? LIMIT ?")
		response, _ = db.Query(strQuery, cursor, categoryId, "%"+search+"%", limit)
	} else if lenSearch == 0 && len(sort) != 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? ORDER BY price " + sort + " LIMIT ?")
		response, _ = db.Query(strQuery, cursor, categoryId, limit)
	} else {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? LIMIT ?")
		response, _ = db.Query(strQuery, cursor, categoryId, limit)
	}

	if isFeature == 1 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND is_feature = 1 LIMIT ?")
		response, _ = db.Query(strQuery, cursor, limit)
	}

	if response == nil {
		return
	}
	var product models.Product
	for response.Next() {
		err := response.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Image, &product.Sale, &product.IsFeature)
		if err != nil {
			panic(err)
		}
		result = append(result, product)
	}
	defer response.Close()
	return result
}

func GetProductById(id int) models.Product {
	var product models.Product
	err := db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Image, &product.Sale, &product.IsFeature)
	if err != nil {
		fmt.Println("Error in GetAllProductById()")
	}
	return product
}



func CreateProduct(product models.Product) (err error) {
	strQuery := "INSERT INTO products(name, description, price, category_id, image, sale, is_feature) VALUES (?,?,?,?,?,?,?)"
	result, err := db.Exec(strQuery, product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature)
	fmt.Println(result)
	return err
}

func UpdateProduct(product models.Product) (err error) {
	
	strQuery, err := db.Prepare("UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, image = ?, sale = ?, is_feature = ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature, product.Id)
	return err
}

func DeleteProductById(id int) (str string, err error) {
	strQuery := "DELETE FROM products WHERE id = ?"
	result, err := db.Exec(strQuery, id)
	fmt.Println(result)
	return "dccm", err
}
