package repository

import (
	"bt/project/models"
	"database/sql"
)

type CategoryRepository struct {
	Db *sql.DB
}

func (cr *CategoryRepository) GetAllCategories() (result []models.Category) {
	response, _ := cr.Db.Query("SELECT * FROM categories")
	defer response.Close()
	var category models.Category
	for response.Next() {
		err := response.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		result = append(result, category)
	}
	return result
}
