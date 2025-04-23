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
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMemoHandler_CreateMemo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		di.UserUsecase.EXPECT().GetUser(gomock.Any(), uint64(1)).Return(&model.User{
			ID:    1,
			Name:  "Tester",
			Email: "test@example.com",
		}, nil)

		di.MemoUsecase.EXPECT().
			CreateMemo(gomock.Any(), gomock.AssignableToTypeOf(&model.Memo{})).
			Return(nil)

		memo := &model.Memo{
			UserID:  1,
			Title:   "新規メモ",
			Content: nil,
			Active:  model.BoolFlagTrue,
		}

		body, _ := json.Marshal(memo)
		req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Set("userID", uint64(1))

		err := di.MemoHandler.CreateMemo(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"title":"新規メモ"`)
		assert.Contains(t, rec.Body.String(), `"active":1`)
	})

	t.Run("invalid request body", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewReader([]byte("invalid json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		_ = di.MemoHandler.CreateMemo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid request body")
	})

	t.Run("validation error", func(t *testing.T) {
		di := testutil.NewTestHandlerDI(t)
		defer di.Cleanup()

		di.UserUsecase.EXPECT().GetUser(gomock.Any(), uint64(10)).Return(nil, nil)

		memo := &model.Memo{
			UserID: 10,
			Title:  "",
		}
		body, _ := json.Marshal(memo)
		req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Set("userID", uint64(10))

		_ = di.MemoHandler.CreateMemo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "title")
	})
}
