package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Contacts     []Contact
	Integrations []Integration

	ID           uint64
	Subdomain    string `gorm:"size:255"`
	AuthCode     string
	AccessToken  string
	RefreshToken string
	Expires      uint64
	UnisenderKey string `gorm:"size:255"`

	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}

type ListAccountFilter struct {
	Page             int
	Limit            int
	NeedRefresh      bool
	AmoAuthorized    *bool
	JoinIntegrations bool
}
