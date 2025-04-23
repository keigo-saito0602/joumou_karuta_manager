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

type MemoHandler struct {
	memoUsecase   usecase.MemoUsecase
	memoValidator *validation.MemoValidator
}

func NewMemoHandler(u usecase.MemoUsecase, v *validation.MemoValidator) *MemoHandler {
	return &MemoHandler{
		memoUsecase:   u,
		memoValidator: v,
	}
}

// ListMemos godoc
// @Summary Get all memos
// @Description ユーザーの一覧を取得する
// @Tags memos
// @Accept json
// @Produce json
// @Success 200 {array} model.Memo
// @Failure 500 {object} map[string]string
// @Router /memos [get]
func (h *MemoHandler) ListMemos(c echo.Context) error {
	memos, err := h.memoUsecase.ListMemos(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, memos)
}

// GetMemo godoc
// @Summary Get memo by ID
// @Description IDでユーザーを取得する
// @Tags memos
// @Accept json
// @Produce json
// @Param id path int true "Memo ID"
// @Success 200 {object} model.Memo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memos/{id} [get]
func (h *MemoHandler) GetMemo(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	memo, err := h.memoUsecase.GetMemo(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, memo)
}

// CreateMemo godoc
// @Summary Create a new memo
// @Description 新しいユーザーを作成する
// @Tags memos
// @Accept json
// @Produce json
// @Param memo body model.Memo true "New memo"
// @Success 201 {object} model.Memo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memos [post]
func (h *MemoHandler) CreateMemo(c echo.Context) error {
	var memo model.Memo
	if err := c.Bind(&memo); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.memoValidator.ValidateCreate(c.Request().Context(), &memo); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	if err := h.memoUsecase.CreateMemo(c.Request().Context(), &memo); err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusCreated, memo)
}

// UpdateMemo godoc
// @Summary Update memo by ID
// @Description 指定されたIDのユーザー情報を更新します
// @Tags memos
// @Accept json
// @Produce json
// @Param id path int true "Memo ID"
// @Param memo body model.Memo true "Updated memo"
// @Success 200 {object} model.Memo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memos/{id} [put]
func (h *MemoHandler) UpdateMemo(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	var memo model.Memo
	if err := c.Bind(&memo); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "invalid request body")
	}
	memo.ID = id

	if err := h.memoValidator.ValidateUpdate(c.Request().Context(), &memo); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	err = h.memoUsecase.UpdateMemo(c.Request().Context(), &memo)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, memo)
}

// DeleteMemo godoc
// @Summary Delete memo by ID
// @Description 指定されたIDのユーザーを削除します
// @Tags memos
// @Accept json
// @Produce json
// @Param id path int true "Memo ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memos/{id} [delete]
func (h *MemoHandler) DeleteMemo(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	err = h.memoUsecase.DeleteMemo(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
