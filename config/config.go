package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Conexão DB
func GetDB() (db *sql.DB, err error) {

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/produtos")

	return
}
