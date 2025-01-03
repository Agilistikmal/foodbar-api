package repository

import (
	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Save(auth *model.Auth) (*model.Auth, error) {
	err := r.db.Save(&auth).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *AuthRepository) Find(phone string) (*model.Auth, error) {
	var auth *model.Auth
	err := r.db.Take(&auth, "phone = ?", phone).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}
