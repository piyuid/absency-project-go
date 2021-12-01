package response

import (
	"absency/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID             uint64 `gorm:"primaryKey"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Pin            uint32 `json:"pin"`
	Previllage     uint32 `json:"previllage"`
	PrevillageUser string `json:"previllageUser"`
	FingerId       string `json:"fingerid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain users.Users) UserResponse {
	return UserResponse{
		ID:             domain.ID,
		Username:       domain.Username,
		Password:       domain.Password,
		Email:          domain.Email,
		Pin:            domain.Pin,
		Previllage:     domain.Previllage,
		PrevillageUser: domain.PrevillageUser,
		FingerId:       domain.FingerId,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
	}
}
