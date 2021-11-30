package position

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint64 `gorm:"primaryKey"`
	CodePosition uint8  `json:"CodePosition"`
	NamePosition string `json:"NamePosition"`
	DescPosition string `json:"DescPosition"`
	Hierarchy    string `json:"Hierarchy"`
	HierarchyId  uint8  `json:"HierarchyId"`
	StaffChild   uint8  `json:"StaffChild"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
