package request

import "absency/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() *users.Users {
	return &users.Users{
		Email:    user.Email,
		Password: user.Password,
	}
}
