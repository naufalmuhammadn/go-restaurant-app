package menu

import (
	"gorm.io/gorm"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model/constant"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

// Signature, GetMenu merupakan methode dari m (menuRepo)
func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem

	if err := m.db.Where(model.MenuItem{Type: constant.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}