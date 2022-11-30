package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	authUsecase "go-clean-api/auth/usecase"
	"go-clean-api/entity"
	"go-clean-api/entity/mocks"
	_userMiddleware "go-clean-api/user/middleware"
	"go-clean-api/util"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var (
	mockConfig = util.Config{
		JwtSecretKey: "secret",
	}
	testDatas = []struct {
		name    string
		request entity.User
		errRes  bool
		message string
	}{
		{
			name: "Test User Allowed",
			request: entity.User{
				Username: "Admin",
				Password: "123123",
				Role:     "administrator",
				Email:    "test1@test.com",
				Id:       1,
				Name:     "Administrator",
			},
			errRes:  false,
			message: "Return must be success",
		},
		{
			name: "Test User Not Allowed",
			request: entity.User{
				Username: "User",
				Password: "123123",
				Role:     "user",
				Email:    "test2@test.com",
				Id:       2,
				Name:     "User",
			},
			errRes:  true,
			message: "Return must be failed",
		},
	}
)

func BuildMockTokenInterface(user *entity.User, loginUc entity.AuthUCaseInterface) *jwt.Token {
	tokenStr := loginUc.GenerateToken(user)
	mockTokenInterface, _ := jwt.ParseWithClaims(tokenStr, &entity.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mockConfig.JwtSecretKey), nil
		})
	return mockTokenInterface
}

func createContextHttpTest(e *echo.Echo) echo.Context {
	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	return c
}

func TestAllowedUserRole(t *testing.T) {
	mockUserRepo := new(mocks.UserRepoInterface)

	loginUc := authUsecase.NewLoginUseCase(mockUserRepo, &mockConfig)

	var wg sync.WaitGroup

	wg.Add(len(testDatas))

	for _, test := range testDatas {
		test := test
		go func() {

			t.Run(test.name, func(t *testing.T) {
				e := echo.New()
				mockTokenInterface := BuildMockTokenInterface(&test.request, loginUc)

				c := createContextHttpTest(e)
				c.Set("jwtInfo", (mockTokenInterface).Claims.(*entity.JwtCustomClaims))

				mUser := _userMiddleware.InitUserMiddleware()
				h := mUser.OnlyAdmin(func(c echo.Context) error {
					return c.NoContent(http.StatusOK)
				})

				err := h(c)
				if err == nil {
					assert.Equal(t, test.errRes, false, test.message)
				} else {
					assert.Equal(t, test.errRes, true, test.message)
				}
			})

			wg.Done()
		}()
	}
	wg.Wait()
}
