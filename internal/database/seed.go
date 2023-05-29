// Seed : menginisialisasi database dengan data awal

package database

import(
	"gorm.io/gorm"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model/constant"
)

func seedDB( db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:        "Indomie Mi Goreng",
			Description: "Fried Indomie noodles, seasoned with a unique blend of spices and seasonings, served hot and ready to eat",
			OrderCode:   "30042301",
			Price:       1.99,
			Type:        constant.MenuTypeFood,
		},
		{
			Name:        "Indomie Kuah Kari Ayam",
			Description: "Indomie noodles served in a hot chicken curry broth, seasoned with a unique blend of spices and seasonings",
			OrderCode:   "30042302",
			Price:       1.99,
			Type:        constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:        "Kopi Aren",
			Description: "A traditional Indonesian coffee drink, made with palm sugar and coconut milk",
			OrderCode:   "30042303",
			Price:       2.99,
			Type:        constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}