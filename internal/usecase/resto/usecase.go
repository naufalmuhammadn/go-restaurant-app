package resto

import "github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}