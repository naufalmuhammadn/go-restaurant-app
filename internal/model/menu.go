package model

type MenuType string

type MenuItem struct {
	Name        string
	Description string
	OrderCode   string
	Price       float64
	Type        MenuType
}
