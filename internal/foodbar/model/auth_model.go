package model

type Auth struct {
	Phone string `gorm:"primaryKey" json:"phone,omitempty"`
	OTP   string `json:"otp,omitempty"`
}
