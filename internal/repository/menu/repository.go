package menu

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}