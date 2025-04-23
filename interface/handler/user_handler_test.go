package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"github.com/keigo-saito0602/joumou_karuta_manager/testutil"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		plainPassword := "password123"
		hash, _ := util.HashPassword(plainPassword)

		di.UserUsecase.EXPECT().GetByEmail(gomock.Any(), "test@example.com").Return(&model.User{
			ID:       1,
			Email:    "test@example.com",
			Password: hash,
			Role:     "user",
		}, nil)

		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": plainPassword,
		}
		reqJSON, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		if assert.NoError(t, di.UserHandler.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var res map[string]string
			err := json.Unmarshal(rec.Body.Bytes(), &res)
			assert.NoError(t, err)
			assert.Contains(t, res, "token")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte("{badjson")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		_ = di.UserHandler.Login(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "不正なリクエストです")
	})

	t.Run("unauthorized", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		di.UserUsecase.EXPECT().GetByEmail(gomock.Any(), "wrong@example.com").Return(&model.User{
			Password: "$2a$10$invalidhash",
		}, nil)

		reqBody := map[string]string{
			"email":    "wrong@example.com",
			"password": "wrongpass",
		}
		reqJSON, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		_ = di.UserHandler.Login(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "メールアドレスまたはパスワードが間違っています")
	})
}
