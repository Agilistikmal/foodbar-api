package repository

import (
	"encoding/csv"
	"os"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) ReadFromCSV() ([]*model.Product, error) {
	file, err := os.Open("./data/products.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var products []*model.Product

	for _, record := range records {
		product := &model.Product{
			Barcode:     record[0],
			Name:        record[1],
			Certificate: record[2],
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) ConvertFromCSV() error {
	products, err := r.ReadFromCSV()
	if err != nil {
		return err
	}

	max := 1000
	for i := 0; i < len(products); i += max {
		m := i + max
		if len(products) < m {
			m = len(products)
		}
		part := products[i:m]
		err := r.db.Save(&part).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ProductRepository) Save(product *model.Product) error {
	return r.db.Save(&product).Error
}

func (r *ProductRepository) Find(barcode string) (*model.Product, error) {
	var product *model.Product
	err := r.db.Take(&product, "barcode = ?", barcode).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
