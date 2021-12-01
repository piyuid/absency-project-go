package users

import (
	_middleware "absency/app/middleware"
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		// jwt:  configJWT,
	}
}

func (usecase *UserUseCase) Login(domain Users, ctx context.Context) (Users, error) {
	if domain.Email == "" {
		return Users{}, errors.New("Email Empty")
	}
	if domain.Password == "" {
		return Users{}, errors.New("Password Empty")
	}
	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Users{}, err
	}

	// user.Token = usecase.jwt.GenererateToken(user.Id)

	return user, nil
}

func (usecase *UserUseCase) GetAllUsers(ctx context.Context) ([]Users, error) {
	return []Users{}, nil
}

// 5, 7 => 12
func Addition(a, b int) int {
	sum := a + b
	if sum < 0 {
		sum = 0
	}
	return sum
}
