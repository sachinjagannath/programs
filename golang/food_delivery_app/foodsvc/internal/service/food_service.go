package service

import (
	"context"
	"sync"
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

// GetMenuItemSummary: performs multiple DB calls concurrently:
// - menu item
// - restaurant
// - "supplier" (simulated) and price formatting
func (s *Service) GetMenuItemSummary(ctx context.Context, menuID int64) (*models.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// channels
	menuCh := make(chan *models.MenuItem, 1)
	restCh := make(chan *models.Restaurant, 1)
	errCh := make(chan error, 2)

	var wg sync.WaitGroup

	wg.Add(2)

	//fetch menu item
	go func() {
		defer wg.Done()
		m, err := s.Repo.GetMenuItem(ctx, menuID)
		if err != nil {
			errCh <- err
			return
		}
		menuCh <- m
	}()

	//fetch Restaurant(needs menu first ; fetch menuId -> restaurantID in seperate go routine pattern)
}
