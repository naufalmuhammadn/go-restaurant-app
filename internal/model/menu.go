package model

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model/constant"
)

type MenuItem struct {
	Name        string
	Description string
	OrderCode   string
	Price       float64
	Type        constant.MenuType
}