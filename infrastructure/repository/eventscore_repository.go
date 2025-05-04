package repository

import (
	"context"
	"log"

	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"gorm.io/gorm"
)

type EventScoreRepository interface {
	CreateEventScore(ctx context.Context, eventScore *model.EventScoreForCreate) error
	GetEventScoreWithRank(ctx context.Context, id uint64) (*model.EventScore, error)
	ListEventScores(ctx context.Context) ([]*model.EventScore, error)
	DeleteAllEventScores(ctx context.Context) error
}

type eventScoreRepository struct {
	db *gorm.DB
}

func NewEventScoreRepository(db *gorm.DB) EventScoreRepository {
	return &eventScoreRepository{db: db}
}

func (r *eventScoreRepository) CreateEventScore(ctx context.Context, eventScore *model.EventScoreForCreate) error {
	err := r.db.
		Table("event_scores").
		Create(eventScore).
		Error
	if err != nil {
		log.Printf("[EventScoreRepository][CreateEventScore] Failed to create eventScore: %v", err)
	}
	return err
}

func (r *eventScoreRepository) GetEventScoreWithRank(ctx context.Context, id uint64) (*model.EventScore, error) {
	var e model.EventScore
	// 対象レコードをIDで取得
	if err := r.db.
		Table("event_scores").
		First(&e, id).
		Error; err != nil {
		return nil, err
	}

	// 取得した対象レコードの得点より上位の件数をカウントする。 rank = count + 1
	var cnt int64
	if err := r.db.
		Table("event_scores").
		Where(`
			score > ? OR
			(score = ? AND cards_taken > ?) OR
			(score = ? AND cards_taken = ? AND fault_count < ?) OR
			(score = ? AND cards_taken = ? AND fault_count = ? AND created_at < ?)
		`,
			e.Score, e.Score, e.CardsTaken,
			e.Score, e.Score, e.FaultCount,
			e.Score, e.Score, e.FaultCount, e.CreatedAt,
		).
		Count(&cnt).
		Error; err != nil {
		return nil, err
	}
	e.Rank = int(cnt) + 1
	return &e, nil
}

func (r *eventScoreRepository) ListEventScores(ctx context.Context) ([]*model.EventScore, error) {
	var tmp []model.EventScore
	if err := r.db.Raw(`
			SELECT
					id, name, feeling, score, cards_taken, fault_count, created_at
			FROM event_scores
			ORDER BY
					score DESC,
					cards_taken DESC,
					fault_count ASC,
					created_at ASC
	`).Scan(&tmp).Error; err != nil {
		return nil, err
	}
	out := make([]*model.EventScore, len(tmp))
	for i := range tmp {
		out[i] = &tmp[i]
	}
	return out, nil
}

func (r *eventScoreRepository) DeleteAllEventScores(ctx context.Context) error {
	log := logger.FromContext(ctx)
	err := r.db.
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&model.EventScore{}).
		Error
	if err != nil {
		log.Errorf("failed to delete eventScores: %v", err)
	}
	return err
}
