package handler

import (
	"net/http"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"github.com/keigo-saito0602/joumou_karuta_manager/usecase"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"github.com/labstack/echo/v4"
)

type CardHandler struct {
	cardUsecase usecase.CardUsecase
}

func NewCardHandler(u usecase.CardUsecase) *CardHandler {
	return &CardHandler{
		cardUsecase: u,
	}
}

// ListCards godoc
// @Summary Get all cards
// @Description ユーザーの一覧を取得する
// @Tags cards
// @Accept json
// @Produce json
// @Success 200 {array} model.Card
// @Failure 500 {object} map[string]string
// @Router /cards [get]
func (h *CardHandler) ListCards(c echo.Context) error {
	cards, err := h.cardUsecase.ListCards(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, cards)
}

// GetCard godoc
// @Summary Get card by ID
// @Description IDでユーザーを取得する
// @Tags cards
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} model.Card
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/{id} [get]
func (h *CardHandler) GetCard(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	card, err := h.cardUsecase.GetCard(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, card)
}

// ShuffleCards godoc
// @Summary Get all cards (shuffled)
// @Description すべてのカードをランダムな順番で取得する
// @Tags cards
// @Accept json
// @Produce json
// @Success 200 {array} model.Card
// @Failure 500 {object} map[string]string
// @Router /cards/shuffle [get]
func (h *CardHandler) ShuffleCards(c echo.Context) error {
	cards, err := h.cardUsecase.ListShuffledCards(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, cards)
}

// ListCardsByInitial godoc
// @Summary Get cards by initial kana
// @Description 指定した五十音で始まるカード一覧を取得する
// @Tags cards
// @Accept json
// @Produce json
// @Param initial query string true "Initial kana (e.g., 'あ')"
// @Success 200 {array} model.Card
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/by-initial [get]
func (h *CardHandler) ListCardsByInitial(c echo.Context) error {
	initial := c.QueryParam("initial")
	if initial == "" {
		return util.ErrorJSON(c, http.StatusBadRequest, "initial is required")
	}

	cards, err := h.cardUsecase.ListCardsBySyllabary(
		c.Request().Context(),
		model.Syllabary(initial),
	)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, cards)
}
