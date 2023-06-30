package model

import (
	"gorm.io/gorm"
)

type Integration struct {
	gorm.Model

	ID        uint64
	OuterID   string `gorm:"size:255"`
	AccountID uint64
	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}
