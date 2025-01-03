package main

import (
	"log"
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/delivery/rest"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/delivery/rest/route"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/repository"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/service"
	"github.com/agilistikmal/foodbar-api/internal/infrastructure/config"
	"github.com/agilistikmal/foodbar-api/internal/infrastructure/database"
)

func main() {
	config.NewConfig()

	db := database.NewDatabase()

	halalRepository := repository.NewHalalRepository(db)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, halalRepository)
	productHandler := rest.NewProductHandler(productService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := rest.NewAuthHandler(authService)

	routes := route.NewRoutes(productHandler, authHandler)
	routes.Init()

	log.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", routes.Mux)
}
