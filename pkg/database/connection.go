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
	fmt.Printf("DSN -  %s", mysqlconn)
	fmt.Printf("MYSQL_USER -  %s", os.Getenv("MYSQL_USER"))
	fmt.Printf("MYSQL_PASSWORD -  %s", os.Getenv("MYSQL_PASSWORD"))
	fmt.Printf("MYSQL_DATABASE -  %s", os.Getenv("MYSQL_DATABASE"))
	// conn, err := sql.Open("mysql", mysqlconn)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("MYSQL_USER"),
	// 	os.Getenv("MYSQL_PASSWORD"),
	// 	os.Getenv("MYSQL_HOST"),
	// 	os.Getenv("MYSQL_PORT"),
	// 	os.Getenv("MYSQL_DB_NAME"),
	// )

	// fmt.Printf("DSN -  %s", dsn)
	// fmt.Printf("MYSQL_USER -  %s", os.Getenv("MYSQL_USER"))
	// fmt.Printf("MYSQL_PASSWORD -  %s", os.Getenv("MYSQL_PASSWORD"))
	// fmt.Printf("MYSQL_DB_NAME -  %s", os.Getenv("MYSQL_DB_NAME"))

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
