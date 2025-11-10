package models

type MenuItem struct {
	ID           int64
	RestaurantID int64
	Name         string
	Price        float64
	Available    bool
}
