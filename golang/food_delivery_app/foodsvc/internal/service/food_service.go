package service

import (
	"context"
	"fmt"
	"time"

	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/internal/models"
	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/repo"
)

type Service struct {
	Repo *repo.MySQLRepo
}

func NewService(r *repo.MySQLRepo) *Service {
	return &Service{Repo: r}
}

// GetMenuItemSummary - menu item
// - restaurant
// - "supplier" (simulated) and price formatting
func (s *Service) GetMenuItemSummary(ctx context.Context, menuID int64) (*models.MenuItem, string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	menuCh := make(chan *models.MenuItem, 1)
	restCh := make(chan *models.Restaurant, 1)
	errCh := make(chan error, 2)

	// Fetch menu item first
	go func() {
		m, err := s.Repo.GetMenuItem(ctx, menuID)
		if err != nil {
			errCh <- err
			return
		}
		menuCh <- m

		// Once menu is fetched, fetch restaurant
		r, err := s.Repo.GetRestaurant(ctx, m.RestaurantID)
		if err != nil {
			errCh <- err
			return
		}
		restCh <- r
	}()

	var menu *models.MenuItem
	var rest *models.Restaurant

	for {
		select {
		case m := <-menuCh:
			menu = m
		case r := <-restCh:
			rest = r
		case err := <-errCh:
			return nil, "", "", err
		case <-ctx.Done():
			return nil, "", "", fmt.Errorf("timeout fetching menu summary")
		}

		if menu != nil && rest != nil {
			break
		}
	}

	supplier := "Supplier OK"
	_ = fmt.Sprintf("₹%.2f", menu.Price)
	return menu, rest.Name, supplier, nil
}

// BulkUpdateAvailability uses a worker pool with concurrency limit
func (s *Service) BulkUpdateAvailability(ctx context.Context, ids []int64, available bool, concurrency int) (int, int) {
	if concurrency <= 0 {
		concurrency = 5
	}

	type job struct{ id int64 }

	jobs := make(chan job, len(ids))
	results := make(chan error, len(ids))

	// spawn workers
	for w := 0; w < concurrency; w++ {
		go func() {
			for j := range jobs {
				//each update has its own context
				cctx, cancel := context.WithTimeout(ctx, 2*time.Second)
				err := s.Repo.UpdateAvailability(cctx, j.id, available)
				cancel()
				results <- err
			}
		}()
	}
	for _, id := range ids {
		jobs <- job{id: id}
	}
	close(jobs)

	//collect results
	success := 0
	failed := 0
	for i := 0; i < len(ids); i++ {
		err := <-results
		if err != nil {
			failed++
		} else {
			success++
		}
	}
	return success, failed
}

// CreateOrder simply ensures menu exists and inserts order

func (s *Service) CreateOrder(ctx context.Context, menuItemID int64, customer string) (int64, error) {
	//get menu item
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	m, err := s.Repo.GetMenuItem(ctx, menuItemID)
	if err != nil {
		return 0, err
	}
	// create order
	return s.Repo.CreateOrder(ctx, menuItemID, m.RestaurantID, customer)
}

// GetPendingOrderSummary aggregates count per restaurant concurrently (fanout)
func (s *Service) GetPendingOrderSummary(ctx context.Context) (map[int64]int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.Repo.GetPendingOrdersCountByRestaurant(ctx)
}
