package rest

import (
	"log"
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/service"
	"github.com/agilistikmal/foodbar-api/internal/pkg"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) Find(w http.ResponseWriter, r *http.Request) {
	barcode := r.PathValue("barcode")
	log.Println(barcode)
	product, err := h.service.Find(barcode)
	if err != nil {
		if err.Error() == "record not found" {
			pkg.SendError(w, http.StatusNotFound, err.Error())
		} else {
			pkg.SendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	pkg.SendSuccess(w, product)
}
