package resto

import (
	"github.com/google/uuid"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/model/constant"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/repository/menu"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/repository/order"
)

type restoUsecase struct {
	menuRepo menu.Repository
	orderRepo order.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
		orderRepo: orderRepo,
	}
}

func (r *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUsecase)Order(request model.OrderMenuRequest) (model.Order, error) {
	// Membuat slice make(tipe data, ukuran)
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID: uuid.New().String(),
			OrderCode: orderProduct.OrderCode,
			Quantity: orderProduct.Quantity,
			TotalPrice: int64(menuData.Price) * int64(orderProduct.Quantity),
			Status: constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID: uuid.New().String(),
		Status: constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
	}

	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createOrderData, nil
}

func (r *restoUsecase)GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}