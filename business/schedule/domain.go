package business

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID uint64 `json:"id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
