package handler

import (
	"net/http"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"github.com/keigo-saito0602/joumou_karuta_manager/usecase"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"github.com/keigo-saito0602/joumou_karuta_manager/validation"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase   usecase.UserUsecase
	userValidator *validation.UserValidator
}

func NewUserHandler(usecase usecase.UserUsecase, validation *validation.UserValidator) *UserHandler {
	return &UserHandler{
		userUsecase:   usecase,
		userValidator: validation,
	}
}

// ListUsers godoc
// @Summary Get all users
// @Description ユーザーの一覧を取得する
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.UserResponse
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) ListUsers(c echo.Context) error {
	users, err := h.userUsecase.ListUsers(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get user by ID
// @Description IDでユーザーを取得する
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.userUsecase.GetUser(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description 新しいユーザーを作成する
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "New user"
// @Success 201 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.userValidator.ValidateCreate(c.Request().Context(), &user); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	created, err := h.userUsecase.CreateUser(c.Request().Context(), &user)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusCreated, created)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description 指定されたIDのユーザー情報を更新します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "Updated user"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "invalid request body")
	}
	user.ID = id

	if err := h.userValidator.ValidateUpdate(c.Request().Context(), &user); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	updated, err := h.userUsecase.UpdateUser(c.Request().Context(), &user)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description 指定されたIDのユーザーを削除します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	err = h.userUsecase.DeleteUser(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// Login godoc
// @Summary ログイン
// @Description メールアドレスとパスワードでログイン
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "ログイン情報"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "不正なリクエストです")
	}

	resp, err := h.userUsecase.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// AdminOnlySetting godoc
// @Summary 管理者専用の設定取得
// @Description 管理者しかアクセスできない設定API
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /auth/admin/settings [get]
func (h *UserHandler) AdminOnlySetting(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "管理者用の設定にアクセスしました！",
	})
}
