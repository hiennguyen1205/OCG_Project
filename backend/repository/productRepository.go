package repository

import (
	"database/sql"
	"fmt"

	"bt/project/models"
)

type ProductRepository struct {
	Db *sql.DB
}

// hàm insert data và database
func (pd *ProductRepository) InitData() {
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
		result, err := pd.Db.Exec(strQuery, category.Name)
		if err != nil {
			panic(err.Error())
		}
		lastCategoryID, _ := result.LastInsertId()
		fmt.Printf("Danh mục mới thêm có ID là: %d\n", lastCategoryID)
	}

	for _, product := range listProducts {
		strQuery := "INSERT INTO products(name, description, price, category_id, image, sale, is_feature) VALUES (?,?,?,?,?,?,?)"
		result, err := pd.Db.Exec(strQuery, product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature)
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
func (pd *ProductRepository) GetAllProducts(limit int, cursor int, search string, sort string, categoryId int, isFeature int, all int) (result []models.Product) {
	var strQuery string
	var response *sql.Rows
	lenSearch := len(search)
	lenSort := len(sort)
	if lenSearch != 0 && lenSort != 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? AND name LIKE ? ORDER BY price " + sort + " LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, categoryId, "%"+search+"%", limit)
	} else if lenSearch != 0 && len(sort) == 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? AND name LIKE ? LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, categoryId, "%"+search+"%", limit)
	} else if lenSearch == 0 && len(sort) != 0 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? ORDER BY price " + sort + " LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, categoryId, limit)
	} else {
		strQuery = ("SELECT * FROM products WHERE id > ? AND category_id = ? LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, categoryId, limit)
	}

	if isFeature == 1 {
		strQuery = ("SELECT * FROM products WHERE id > ? AND is_feature = 1 LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, limit)
	}

	if all == 1 {
		strQuery = ("SELECT * FROM products WHERE id > ? LIMIT ?")
		response, _ = pd.Db.Query(strQuery, cursor, limit)
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

func (pd *ProductRepository) GetProductById(id int) models.Product {
	var product models.Product
	err := pd.Db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Image, &product.Sale, &product.IsFeature)
	if err != nil {
		fmt.Println("Error in GetAllProductById()", err)
	}
	return product
}

func (pd *ProductRepository) CreateProduct(product models.Product) (err error) {
	strQuery := "INSERT INTO products(name, description, price, category_id, image, sale, is_feature) VALUES (?,?,?,?,?,?,?)"
	result, err := pd.Db.Exec(strQuery, product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature)
	fmt.Println(result)
	return err
}

func (pd *ProductRepository) UpdateProduct(product models.Product) (err error) {

	strQuery, err := pd.Db.Prepare("UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, image = ?, sale = ?, is_feature = ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale, product.IsFeature, product.Id)
	return err
}

func (pd *ProductRepository) DeleteProductById(id int) (str string, err error) {
	//xóa foreign key
	strQueryOrderItem := "DELETE FROM order_items WHERE product_id = ?"
	pd.Db.Exec(strQueryOrderItem, id)
	//xóa sản phẩm
	strQuery := "DELETE FROM products WHERE id = ?"
	result, err := pd.Db.Exec(strQuery, id)
	fmt.Println(result)
	return "Xóa thành công!", err
}
