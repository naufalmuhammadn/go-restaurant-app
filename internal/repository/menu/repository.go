package menu

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}