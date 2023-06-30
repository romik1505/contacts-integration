package model

import (
	"gorm.io/gorm"
	"regexp"
)

type Contact struct {
	gorm.Model

	ID            uint64 `gorm:"primaryKey" json:"id,omitempty"`
	AmoID         uint64 `json:"amo_id,omitempty"`
	AccountID     uint64 `json:"account_id,omitempty"`
	Name          string `gorm:"size:255" json:"name,omitempty"`
	Email         string `gorm:"size:255" json:"email,omitempty"`
	Type          string `gorm:"size:255" json:"type,omitempty"`
	ReasonOutSync string `gorm:"size:255"`
	Sync          bool

	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}

var emailReg = regexp.MustCompile("^[\\d\\w]+@[\\d\\w]+\\.[\\d\\w]+")

func (c Contact) Valid() bool {
	return emailReg.MatchString(c.Email)
}
