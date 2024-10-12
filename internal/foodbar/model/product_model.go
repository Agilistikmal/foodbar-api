package model

type Product struct {
	Barcode     string `gorm:"primaryKey" json:"barcode,omitempty"`
	Name        string `json:"name,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}
