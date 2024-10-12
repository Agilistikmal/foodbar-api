package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/repository"
	"github.com/spf13/viper"
)

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Find(barcode string) (*model.Product, error) {
	product, err := s.repository.Find(barcode)
	if err != nil {
		return nil, err
	}

	if product.Certificate == "" {
		halalRequest := &model.HalalRequest{
			NamaProduct: strings.Split(product.Name, " ")[0],
			SecretCode:  viper.GetString("halalmui.secret_code"),
		}

		halalRequestJson, err := json.Marshal(halalRequest)
		if err != nil {
			return nil, err
		}

		resp, err := http.Post(viper.GetString("halalmui.base_url")+"/search_product", "application/json", bytes.NewBuffer(halalRequestJson))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var halalResponse model.HalalResponse
		err = json.NewDecoder(resp.Body).Decode(&halalResponse)
		if err != nil {
			return nil, err
		}

		for _, data := range halalResponse.Data {
			if strings.Contains(strings.ToUpper(data.NamaProdusen), "PT") ||
				strings.Contains(strings.ToUpper(data.NamaProdusen), "COMPANY") ||
				strings.Contains(strings.ToUpper(data.NamaProdusen), "CV") {

				product.Certificate = data.NomorSertifikat
				s.repository.Save(product)
				break
			}
		}

	}

	return product, nil
}
