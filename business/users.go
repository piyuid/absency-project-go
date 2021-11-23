package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Pin            uint32 `json:"pin"`
	Previllage     uint32 `json:"previllage"`
	PrevillageUser string `json:"previllageuser"`
	FingerId       string `json:"fingerid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
