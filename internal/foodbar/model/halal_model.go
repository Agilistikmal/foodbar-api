package model

import "time"

type HalalResponse struct {
	Message string       `json:"message,omitempty"`
	Status  string       `json:"status,omitempty"`
	Total   int          `json:"total,omitempty"`
	Data    []*HalalData `json:"data,omitempty"`
}

type HalalData struct {
	ProductId        int        `gorm:"primaryKey" json:"product_id,omitempty"`
	CustomerId       int        `json:"customer_id,omitempty"`
	ExpireDate       string     `json:"expire_date,omitempty"`
	NamaProduk       string     `json:"nama_produk,omitempty"`
	NamaProdusen     string     `json:"nama_produsen,omitempty"`
	NomorSertifikat  string     `json:"nomor_sertifikat,omitempty"`
	ProductGroup     string     `json:"product_group,omitempty"`
	ProductGroupCode string     `json:"product_group_code,omitempty"`
	UpdatedAt        *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type HalalRequest struct {
	NamaProduct string `json:"nama_product"`
	Produsen    string `json:"produsen"`
	CertifiedNo string `json:"certified_no"`
	ValidEnd    string `json:"valid_end"`
	SecretCode  string `json:"secret_code"`
}
