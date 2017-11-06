package DBConnection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetConnection() *sql.DB{
	db, err := sql.Open("mysql","root:ep1137@tcp(127.0.0.1:3306)/playmaths")
	if err != nil {
		log.Fatal(err)
	}

	return db
}


