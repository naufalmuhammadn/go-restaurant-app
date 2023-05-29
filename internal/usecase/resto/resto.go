package resto

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/repository/menu"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}