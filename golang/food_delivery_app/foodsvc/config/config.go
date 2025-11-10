package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // <- this registers the mysql driver
)

func ConnectMySql() *sql.DB {
	user := "fooduser"
	pass := "foodPass123!"
	host := "127.0.0.1:3306"
	name := "foodsvc"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&parseTime=true",
		user, pass, host, name)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("mysql open: %v ", err)
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	if err := db.Ping(); err != nil {
		log.Fatalf("mysql ping: %v ", err)
	}
	return db
}
