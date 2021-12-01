package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID             uint64       `gorm:"primaryKey"`
	Username       string       `json:"username"`
	Password       string       `json:"password"`
	Email          string       `json:"email"`
	Pin            uint32       `json:"pin"`
	Previllage     uint32       `json:"previllage"`
	PrevillageUser string       `json:"previllageuser"`
	FingerId       string       `json:"fingerid"`
	Identities     []Identity   `gorm:"foreignKey:ID"`
	Positions      []Position   `gorm:"foreignKey:ID"`
	Departmens     []Department `gorm:"foreignKey:ID`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type Identity struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Birthday      time.Time `json:"birthday"`
	Status        string    `json:"status"`
	Gender        string    `json:"gender"`
	Npwp          string    `json:"npwp"`
	Ktp           uint16    `json:"ktp"`
	Bpjs          string    `json:"bpjs"`
	HpEmployee    string    `json:"HpEmployee`
	NameEmergency string    `json:"NameEmergency"`
	HpEmergency   string    `json:"HpEmergency`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Position struct {
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

type Department struct {
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

type UserUsecaseInterface interface {
	Login(domain Users, ctx context.Context) (Users, error)
	GetAllUsers(ctx context.Context) ([]Users, error)
}

type UserRepoInterface interface {
	Login(domain Users, ctx context.Context) (Users, error)
	// GetAllUsers(ctx context.Context) ([]Domain, error)
}
