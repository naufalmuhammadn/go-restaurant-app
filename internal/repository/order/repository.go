package order

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
)

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}
