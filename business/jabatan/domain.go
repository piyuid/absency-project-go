package business

import (
	"time"

	"gorm.io/gorm"
)

type JabatanDomain struct {
	ID          uint64 `json:"id"`
	CodeJabatan uint8  `json:"CodeDepartment"`
	NameJabatan string `json:"NameDepartment"`
	DescJabatan string `json:"DescDepartment"`
	Hirarki     string `json:"Hirarki"`
	HirarkiId   uint8  `json:"HirarkiId"`
	StaffChild  uint8  `json:"StaffChild"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
