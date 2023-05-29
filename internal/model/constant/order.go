package constant

import (
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
)

const (
	OrderStatusProcessed model.OrderStatus = "processed"
	OrderStatusFinished model.OrderStatus = "finished"
	OrderStatusFailed model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinished model.ProductOrderStatus = "finished"
)