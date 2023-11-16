package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	mysqlconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	conn, err := sql.Open("mysql", mysqlconn)

	if err != nil {
		return nil, err

	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
