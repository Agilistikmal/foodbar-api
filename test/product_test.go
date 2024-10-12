package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/repository"
	"github.com/agilistikmal/foodbar-api/internal/infrastructure/config"
	"github.com/agilistikmal/foodbar-api/internal/infrastructure/database"
)

var r *repository.ProductRepository

func init() {
	config.NewConfig()
	db := database.NewDatabase()
	r = repository.NewProductRepository(db)
}

func TestReadProduct(t *testing.T) {
	records, err := r.ReadFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records[1].Barcode)
}

func TestConvertFromCSV(t *testing.T) {
	err := r.ConvertFromCSV()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateProduct(t *testing.T) {
	err := r.Save(&model.Product{
		Barcode: "123",
		Name:    "Name",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindProduct(t *testing.T) {
	product, err := r.Find("089686017755")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)
}
