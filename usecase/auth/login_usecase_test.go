package usecase

import (
	"github.com/keigo-saito0602/JoumouKarutaTyping/domain/entity"
	"github.com/keigo-saito0602/JoumouKarutaTyping/entity/mocks"
	userUsecase "github.com/keigo-saito0602/JoumouKarutaTyping/user/usecase"
	"github.com/keigo-saito0602/JoumouKarutaTyping/util"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var config = util.Config{}

func TestValidateLoginParam(t *testing.T) {
	mockUserRepo := new(mocks.UserRepoInterface)

	loginUc := NewLoginUseCase(mockUserRepo, &config)

	tests := []struct {
		name     string
		request  *entity.Login
		errorRes bool
		message  string
	}{
		{
			name: "Check Right Login Param",
			request: &entity.Login{
				Username: "admin",
				Password: "123123",
			},
			errorRes: false,
			message:  "The Result must be login!!!",
		},
		{
			name: "Check Wrong Login Param",
			request: &entity.Login{
				Username: "admin",
			},
			errorRes: true,
			message:  "Need Password Excpected",
		},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	for _, test := range tests {
		test := test
		go func() {
			t.Run(test.name, func(t *testing.T) {
				err := loginUc.ValidateLogin(test.request)
				if err == nil {
					assert.Equal(t, test.errorRes, false, test.message)
				} else {
					assert.Equal(t, test.errorRes, true, test.message)
				}
			})
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestCheckHashPassword(t *testing.T) {
	timeoutContext := time.Duration(30) * time.Second
	mockUserRepo := new(mocks.UserRepoInterface)

	loginUc := NewLoginUseCase(mockUserRepo, &config)
	userUc := userUsecase.NewUserUseCase(mockUserRepo, timeoutContext)

	initPassword := "123123"

	tests := []struct {
		name     string
		request  string
		expected bool
		message  string
	}{
		{
			name:     "CheckRightPassword",
			request:  "123123",
			expected: true,
			message:  "The Result must be True",
		},
		{
			name:     "CheckWrongPassword",
			request:  "1231234",
			expected: false,
			message:  "The Result must be True",
		},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	for _, test := range tests {
		test := test
		go func() {
			t.Run(test.name, func(t *testing.T) {
				hashPassword, _ := userUc.HashPassword(test.request)

				valid := loginUc.CheckHashPassword(initPassword, hashPassword)

				assert.Equal(t, test.expected, valid, test.message)
			})
			wg.Done()
		}()
	}

	wg.Wait()
}
