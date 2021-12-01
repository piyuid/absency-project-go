package users

import (
	"absency/business/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Login(domain users.Users, ctx context.Context) (users.Users, error) {
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Users{}, errors.New("Email not found")
		}
		return users.Users{}, errors.New("Error in database")
	}
	return userDb.ToDomain(), nil
}

// func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
// 	return []users.Domain{}, nil
// }
