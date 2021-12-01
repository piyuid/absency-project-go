package schedule

import (
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	ID        uint64    `gorm:"primaryKey"`
	Regulars  []Regular `gorm:"foreignKey:ID"`
	Holidays  []Holiday `gorm:"foreignKey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Regular struct {
	ID           uint64    `gorm:"primaryKey"`
	DateRegular  time.Time `json:"DateRegular"`
	InRegular    time.Time `json:"InRegular"`
	OutRegular   time.Time `json:"OutRegular"`
	Category     string    `json:"Category"`
	CategoryDesc string    `json:"CategoryDesc"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Holiday struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
