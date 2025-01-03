package route

import (
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	ProductHandler *rest.ProductHandler
	AuthHandler    *rest.AuthHandler
}

func NewRoutes(productHandler *rest.ProductHandler, authHandler *rest.AuthHandler) *Route {
	return &Route{
		Mux:            http.NewServeMux(),
		ProductHandler: productHandler,
		AuthHandler:    authHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	r.Mux.HandleFunc("GET /product/{barcode}", r.ProductHandler.Find)
	r.Mux.HandleFunc("GET /search/{query}", r.ProductHandler.Search)

	r.Mux.HandleFunc("POST /otp/send/{phone}", r.AuthHandler.SendOTP)
	r.Mux.HandleFunc("POST /otp/verify", r.AuthHandler.VerifyOTP)
}
