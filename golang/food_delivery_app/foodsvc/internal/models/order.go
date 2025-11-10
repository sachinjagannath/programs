package models

import "time"

type Order struct {
	ID           int64
	MenuItemID   int64
	RestaurantID int64
	CustomerName string
	Status       string
	CreatedAt    time.Time
}
