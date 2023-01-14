package connect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	var (
		username = os.Getenv("MYSQL_USERNAME")
		password = os.Getenv("MYSQL_PASSWORD")
		hostname = os.Getenv("MYSQL_HOSTNAME")
		dbname   = os.Getenv("MYSQL_DBNAME")
	)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", dsn())
	// fmt.Println(dsn())
	if err != nil {
		panic("Error when opening DB")
	}
	// fmt.Println("Connect successed")
	return db
}
