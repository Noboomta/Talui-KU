package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

var (
	DB *sql.DB
)

func ConnectMySQL(uri string) {
	fmt.Println(uri)
	db, err := sql.Open("mysql", uri)
	DB = db
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}