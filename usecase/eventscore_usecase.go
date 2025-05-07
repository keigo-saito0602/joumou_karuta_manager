package usecase

import (
	"context"

	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"gorm.io/gorm"
)

type EventScoreUsecase interface {
	CreateEventScore(ctx context.Context, req *model.EventScoreForCreate) error
	GetEventScoreWithRank(ctx context.Context, id uint64) (*model.EventScore, error)
	ListEventScoresWithRank(ctx context.Context) ([]*model.EventScore, error)
	DeleteAllEventScores(ctx context.Context) error
}

type eventScoreUsecase struct {
	db                   *gorm.DB
	eventScoreRepository repository.EventScoreRepository
}

func NewEventScoreUsecase(db *gorm.DB, repo repository.EventScoreRepository) EventScoreUsecase {
	return &eventScoreUsecase{db: db, eventScoreRepository: repo}
}

func (e *eventScoreUsecase) CreateEventScore(ctx context.Context, eventScore *model.EventScoreForCreate) error {
	ctx = dbctx.ToContext(ctx, e.db)
	log := logger.FromContext(ctx)
	log.Infof("CreateEventScore called: %+v", eventScore)

	eventScore.Score = calculateScore(eventScore.CardsTaken, eventScore.FaultCount)

	if err := e.eventScoreRepository.CreateEventScore(ctx, eventScore); err != nil {
		log.Errorf("failed to create eventScores: %v", err)
		return err
	}

	return nil
}

func (e *eventScoreUsecase) GetEventScoreWithRank(ctx context.Context, id uint64) (*model.EventScore, error) {
	ctx = dbctx.ToContext(ctx, e.db)
	log := logger.FromContext(ctx)
	log.Infof("GetEventScore called with ID=%d", id)

	eventScores, err := e.eventScoreRepository.GetEventScoreWithRank(ctx, id)
	if err != nil {
		log.Errorf("failed to get eventScores: %v", err)
	}
	return eventScores, err
}

func (u *eventScoreUsecase) ListEventScoresWithRank(ctx context.Context) ([]*model.EventScore, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListEventScoresWithRank called")

	scores, err := u.eventScoreRepository.ListEventScores(ctx)
	if err != nil {
		log.Errorf("failed to list: %v", err)
		return nil, err
	}

	assignRanks(scores)

	totalUsers := len(scores)
	for _, e := range scores {
		e.TotalUsers = totalUsers
	}

	return scores, nil
}

func (e *eventScoreUsecase) DeleteAllEventScores(ctx context.Context) error {
	ctx = dbctx.ToContext(ctx, e.db)
	log := logger.FromContext(ctx)

	err := e.eventScoreRepository.DeleteAllEventScores(ctx)
	if err != nil {
		log.Errorf("failed to delete eventScores: %v", err)
	}
	return err
}

// calculateScore は cardsTaken, faultCount からスコアを算出して 0 未満を 0 に丸める。
func calculateScore(cardsTaken, faultCount int) int {
	raw := (cardsTaken*model.CardsTakenMultiplier - faultCount*model.FaultCountMultiplier)
	score := raw / model.ScoreDivisor
	if score < 0 {
		return 0
	}
	return score
}

// assignRanks ランニングを付与する。
// 同じスコアには同じ順位を、スコアが変わった行にはその行番号(i+1)を順位としてセットする。
func assignRanks(scores []*model.EventScore) {
	var prevScore int
	rank := 0

	// i==0: 最初の要素。必ず rank = 1 にする。
	for i, e := range scores {
		if i == 0 || e.Score != prevScore {
			rank = i + 1
			prevScore = e.Score
		}
		e.Rank = rank
	}
}
