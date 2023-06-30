package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Contacts     []Contact
	Integrations []Integration

	ID              uint64
	Subdomain       string `gorm:"size:255"`
	AuthCode        string
	AccessToken     string
	RefreshToken    string
	Expires         uint64
	UnisenderKey    string `gorm:"size:255"`
	UnisenderListID uint64

	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}
