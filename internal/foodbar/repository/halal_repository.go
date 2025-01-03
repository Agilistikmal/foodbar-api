package repository

import (
	"strings"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"gorm.io/gorm"
)

type HalalRepository struct {
	db *gorm.DB
}

func NewHalalRepository(db *gorm.DB) *HalalRepository {
	return &HalalRepository{
		db: db,
	}
}

func (r *HalalRepository) Search(query string) ([]*model.HalalData, error) {
	var result []*model.HalalData
	queryLike := "%" + strings.ToLower(query) + "%"
	err := r.db.Find(&result, "LOWER(nama_produk) LIKE ? AND updated_at > now() - interval '24 hours'", queryLike).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *HalalRepository) Save(halalData *model.HalalData) (*model.HalalData, error) {
	err := r.db.Save(&halalData).Error
	if err != nil {
		return nil, err
	}

	return halalData, nil
}
