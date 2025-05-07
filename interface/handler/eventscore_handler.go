package handler

import (
	"fmt"
	"net/http"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"github.com/keigo-saito0602/joumou_karuta_manager/usecase"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"github.com/keigo-saito0602/joumou_karuta_manager/validation"
	"github.com/labstack/echo/v4"
)

type EventScoreHandler struct {
	eventScoreUsecase   usecase.EventScoreUsecase
	eventScoreValidator *validation.EventScoreValidator
}

func NewEventScoreHandler(u usecase.EventScoreUsecase, v *validation.EventScoreValidator) *EventScoreHandler {
	return &EventScoreHandler{
		eventScoreUsecase:   u,
		eventScoreValidator: v,
	}
}

// CreateEventScore godoc
// @Summary Create a new eventScore
// @Description 新しいイベントスコアを作成する
// @Tags eventScores
// @Accept json
// @Produce json
// @Param eventScore body model.EventScore true "New eventScore"
// @Success 201 {object} model.EventScore
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event_scores [post]
func (h *EventScoreHandler) CreateEventScore(c echo.Context) error {
	var eventScore model.EventScoreForCreate
	if err := c.Bind(&eventScore); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, "invalid request body")
	}
	fmt.Println("test")
	if err := h.eventScoreValidator.CreateEventScoreValidator(c.Request().Context(), &eventScore); err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	if err := h.eventScoreUsecase.CreateEventScore(c.Request().Context(), &eventScore); err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusCreated, eventScore)
}

// GetEventScoreWithRank godoc
// @Summary Get eventScore by ID
// @Description IDでイベントスコアを取得し順位を付加して返す
// @Tags eventScores
// @Accept json
// @Produce json
// @Param id path int true "EventScore ID"
// @Success 200 {object} model.EventScore
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event_scores/{id} [get]
func (h *EventScoreHandler) GetEventScoreWithRank(c echo.Context) error {
	idParam := c.Param("id")
	id, err := util.ParseUint64Param(idParam)
	if err != nil {
		return util.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	eventScore, err := h.eventScoreUsecase.GetEventScoreWithRank(c.Request().Context(), id)
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.JSON(http.StatusOK, eventScore)
}

// ListEventScoresWithRank godoc
// @Summary Get all eventScores
// @Description イベントスコアの一覧を取得する（全件ランキング）
// @Tags eventScores
// @Accept json
// @Produce json
// @Success 200 {array} model.EventScore
// @Failure 500 {object} map[string]string
// @Router /event_scores [get]
func (h *EventScoreHandler) ListEventScoresWithRank(c echo.Context) error {
	eventScores, err := h.eventScoreUsecase.ListEventScoresWithRank(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}
	return c.JSON(http.StatusOK, eventScores)
}

// DeleteAllEventScores godoc
// @Summary Delete all eventScores
// @Description イベントスコアをすべて削除します
// @Tags eventScores
// @Accept json
// @Produce json
// @Param id path int true "EventScore ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event_scores [delete]
func (h *EventScoreHandler) DeleteAllEventScores(c echo.Context) error {
	err := h.eventScoreUsecase.DeleteAllEventScores(c.Request().Context())
	if err != nil {
		status := domain.ErrorToHTTPStatus(err)
		return util.ErrorJSON(c, status, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
