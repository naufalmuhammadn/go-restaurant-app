package main

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/database"
	mRepo "github.com/naufalmuhammadn/go-restaurant-app.git/internal/repository/menu"
	rUsecase "github.com/naufalmuhammadn/go-restaurant-app.git/internal/usecase/resto"
	"github.com/naufalmuhammadn/go-restaurant-app.git/internal/delivery/rest"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=12345 dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)

	// Untuk mengakses data
	menuRepo := mRepo.GetRepository(db)

	// Untuk logic, implementasi operasi
	restoUsecase := rUsecase.GetUsecase(menuRepo)

	// Handler request
	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":3000"))
}