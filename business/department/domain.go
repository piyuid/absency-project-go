package department

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint64 `json:"id"`
	CodeDepartment uint8  `json:"CodeDepartment"`
	NameDepartment string `json:"NameDepartment"`
	DescDepartment string `json:"DescDepartment"`
	Parents        string `json:"Parents"`
	ParentsId      uint8  `json:"ParentsId"`
	Child          uint8  `json:"Child"`
	LogoDepartment string `json:"LogoDepartment"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type DepartmentRepository interface {
	GetByID(ctx context.Context, id int64) (Department, error)
}
