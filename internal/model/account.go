package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Contacts     []Contact
	Integrations []Integration

	ID              uint64
	Subdomain       string `gorm:"size:255"`
	AuthCode        sql.NullString
	AccessToken     sql.NullString
	RefreshToken    sql.NullString
	Expires         sql.NullInt64
	UnisenderKey    string `gorm:"size:255"`
	UnisenderListID uint64

	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}
