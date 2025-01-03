package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/service"
	"github.com/agilistikmal/foodbar-api/internal/pkg"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var request *model.Auth
	json.NewDecoder(r.Body).Decode(&request)

	err := h.service.VerifyOTP(request)
	if err != nil {
		pkg.SendError(w, http.StatusNotFound, err.Error())
		return
	}
	pkg.SendSuccess(w, "verified")
}

func (h *AuthHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	phone := r.PathValue("phone")
	err := h.service.SendOTP(phone)
	if err != nil {
		pkg.SendError(w, http.StatusNotFound, err.Error())
		return
	}
	pkg.SendSuccess(w, "sent")
}
