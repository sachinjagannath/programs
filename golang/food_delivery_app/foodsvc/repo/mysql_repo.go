package repo

import (
	"context"
	"database/sql"

	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/internal/models"
)

const (
	getMenuItemQuery        = "SELECT id, restaurant_id, name, price, available FROM menu_items WHERE id = ?"
	getRestaurantQuery      = "SELECT id, name, location FROM restaurants WHERE id = ?"
	updateAvailabilityQuery = "UPDATE menu_items SET available = ? WHERE id = ?"
	createOrderQuery        = "INSERT INTO orders (menu_item_id, restaurant_id, customer_name, status) VALUES (?, ?, ?, ?)"
	pendingOrdersCountQuery = "SELECT restaurant_id, COUNT(*) FROM orders WHERE status = ? GROUP BY restaurant_id"

	// Status constants
	orderStatusPending = "pending"
)

type MySQLRepo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *MySQLRepo {
	return &MySQLRepo{DB: db}
}

func (r *MySQLRepo) GetMenuItem(ctx context.Context, id int64) (*models.MenuItem, error) {
	row := r.DB.QueryRowContext(ctx, getMenuItemQuery, id)
	m := &models.MenuItem{}
	if err := row.Scan(&m.ID, &m.RestaurantID, &m.Name, &m.Price, &m.Available); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *MySQLRepo) GetRestaurant(ctx context.Context, id int64) (*models.Restaurant, error) {
	row := r.DB.QueryRowContext(ctx, getRestaurantQuery, id)
	s := &models.Restaurant{}
	if err := row.Scan(&s.ID, &s.Location, &s.Name); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *MySQLRepo) UpdateAvailability(ctx context.Context, id int64, available bool) error {
	_, err := r.DB.ExecContext(ctx, updateAvailabilityQuery, available, id)
	return err
}

func (r *MySQLRepo) CreateOrder(ctx context.Context, itemID, restaurantID int64, customer string) (int64, error) {
	res, err := r.DB.ExecContext(ctx, createOrderQuery, itemID, restaurantID, customer, orderStatusPending)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

// Pending orders per restaurant
func (r *MySQLRepo) GetPendingOrdersCountByRestaurant(ctx context.Context) (map[int64]int, error) {
	rows, err := r.DB.QueryContext(ctx, pendingOrdersCountQuery, orderStatusPending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[int64]int)
	for rows.Next() {
		var rid int64
		var cnt int
		if err := rows.Scan(&rid, &cnt); err != nil {
			return nil, err
		}
		counts[rid] = cnt
	}
	return counts, nil
}
