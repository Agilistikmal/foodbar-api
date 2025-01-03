package service

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/agilistikmal/foodbar-api/internal/foodbar/repository"
	"github.com/agilistikmal/foodbar-api/internal/pkg"
	"github.com/spf13/viper"
)

type AuthService struct {
	repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) SendOTP(phone string) error {
	code := pkg.RandomString(4)

	body := []byte(fmt.Sprintf(`{
		"session": "default",
		"chatId":  "%s",
		"text":    "%s"
	}`, phone+"@c.us", fmt.Sprintf("OTP: %s", code)))

	endpoint := viper.GetString("waha.base_url") + "/api/sendText"

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("failed to send otp")
	}

	auth := &model.Auth{
		Phone: phone,
		OTP:   code,
	}

	_, err = s.repository.Save(auth)

	return err
}

func (s *AuthService) VerifyOTP(request *model.Auth) error {
	auth, err := s.repository.Find(request.Phone)
	if err != nil {
		return err
	}

	if auth.OTP != request.OTP {
		return fmt.Errorf("invalid otp")
	}

	return nil
}
