package users_test

import (
	"absency/business/users"
	"absency/business/users/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoInterfaceMock mocks.UserRepoInterface
var userUseCaseInterface users.UserUsecaseInterface
var userDataDummyLogin users.Users

func setup() {
	userUseCaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1)
	userDataDummyLogin = users.Users{
		ID:       1,
		Username: "Leo",
		Email:    "leopuji17@gmail.com",
		Password: "apaaja",
		FingerId: "",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Success Login", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.Users"), mock.Anything).Return(userDataDummyLogin, nil).Once()

		var requestLoginDomain = users.Users{
			Email:    "leopuji17@gmail.com",
			Password: "apaaja",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, domain)
	})

	t.Run("Login with Email Empty", func(t *testing.T) {
		var requestLoginDomain = users.Users{
			Email:    "",
			Password: "apaaja",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Email entered is Blank", err.Error())
		assert.Equal(t, users.Users{}, domain)
	})

	t.Run("Your Password entered is Blank", func(t *testing.T) {
		var requestLoginDomain = users.Users{
			Email:    "leopuji17@gmail.com",
			Password: "",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Password no Fill", err.Error())
		assert.Equal(t, users.Users{}, domain)
	})

	t.Run("Email not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.Domain"), mock.Anything).Return(users.Users{}, errors.New("Email not found")).Once()

		var requestLoginDomain = users.Users{
			Email:    "leopuji17@gmail.com",
			Password: "apaaja",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, errors.New("Email not found"), err)
		assert.Equal(t, users.Users{}, domain)
	})

}
