package route

import (
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	ProductHandler *rest.ProductHandler
}

func NewRoutes(productHandler *rest.ProductHandler) *Route {
	return &Route{
		Mux:            http.NewServeMux(),
		ProductHandler: productHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	r.Mux.HandleFunc("GET /product/{barcode}", r.ProductHandler.Find)
}
