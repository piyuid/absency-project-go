package users

import (
	"absency/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint64 `gorm:"primaryKey"`
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

func (user User) ToDomain() users.Users {
	return users.Users{
		ID:             user.ID,
		Username:       user.Username,
		Password:       user.Password,
		Email:          user.Email,
		Pin:            user.Pin,
		Previllage:     user.Previllage,
		PrevillageUser: user.PrevillageUser,
		FingerId:       user.FingerId,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

func FromDomain(domain users.Users) User {
	return User{
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
