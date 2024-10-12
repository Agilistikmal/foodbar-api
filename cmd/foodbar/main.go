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
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := rest.NewProductHandler(productService)

	routes := route.NewRoutes(productHandler)
	routes.Init()

	log.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", routes.Mux)
}
