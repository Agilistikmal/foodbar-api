package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/repository"
	"github.com/k3a/html2text"
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
		resp, err := http.Get(fmt.Sprintf("http://halal.addi.is.its.ac.id/apiv2?q=%s&result=5", url.QueryEscape(product.Name)))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		jsonString := html2text.HTML2Text(string(b))

		fmt.Println(jsonString)

		var halalResponse model.HalalResponse
		err = json.Unmarshal([]byte(jsonString), &halalResponse)
		if err != nil {
			return nil, err
		}

		for _, data := range halalResponse.EntityData {
			if strings.Contains(data.Label, strings.Split(product.Name, " ")[0]) {
				product.Certificate = data.Atribute.Certificate
				s.repository.Save(product)
				break
			}
		}
	}

	return product, nil
}
