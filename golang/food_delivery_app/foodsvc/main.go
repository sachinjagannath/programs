package main

import (
	"fmt"
	"log"

	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/config"
)

func main() {
	db := config.ConnectMySql()
	defer db.Close()
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Println("Connected to MySQL! Version:", version)
}
