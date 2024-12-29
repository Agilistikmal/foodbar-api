package model

type Product struct {
	Barcode     string `gorm:"primaryKey" json:"barcode,omitempty"`
	Name        string `json:"name,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}

type ProductWithHalalData struct {
	Product   *Product   `json:"product,omitempty"`
	HalalData *HalalData `json:"halal_data,omitempty"`
}
