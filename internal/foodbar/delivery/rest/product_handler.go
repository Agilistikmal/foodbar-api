package rest

import (
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

func (h *ProductHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.PathValue("query")
	products, err := h.service.Search(query)
	if err != nil {
		if err.Error() == "record not found" {
			pkg.SendError(w, http.StatusNotFound, err.Error())
		} else {
			pkg.SendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	pkg.SendSuccess(w, products)
}
