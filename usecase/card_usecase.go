package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"gorm.io/gorm"
)

type CardUsecase interface {
	GetCard(ctx context.Context, id uint64) (*model.Card, error)
	ListCards(ctx context.Context) ([]model.Card, error)
	ListShuffledCards(ctx context.Context) ([]model.Card, error)
	ListCardsBySyllabary(ctx context.Context, syllabary model.Syllabary) ([]model.Card, error)
}

// usecase/card_usecase.go
type cardUsecase struct {
	db             *gorm.DB
	cardRepository repository.CardRepository
}

func NewCardUsecase(db *gorm.DB, cardRepo repository.CardRepository) CardUsecase {
	return &cardUsecase{
		db:             db,
		cardRepository: cardRepo,
	}
}

func (u *cardUsecase) GetCard(ctx context.Context, id uint64) (*model.Card, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("GetCard called with ID=%d", id)

	card, err := u.cardRepository.GetCard(ctx, id)
	if err != nil {
		log.Errorf("failed to get card: %v", err)
	}
	return card, err
}

func (u *cardUsecase) ListCards(ctx context.Context) ([]model.Card, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListCards called")

	cards, err := u.cardRepository.ListCards(ctx)
	if err != nil {
		log.Errorf("failed to list cards: %v", err)
	}
	return cards, err
}

func (u *cardUsecase) ListShuffledCards(ctx context.Context) ([]model.Card, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListShuffledCards called")

	cards, err := u.cardRepository.ListCards(ctx)
	if err != nil {
		log.Errorf("failed to list cards for shuffle: %v", err)
		return nil, err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards, nil
}

func (u *cardUsecase) ListCardsBySyllabary(ctx context.Context, syllabary model.Syllabary) ([]model.Card, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListCardsBySyllabary called")

	initials := model.GetInitialsBySyllabary(syllabary)
	if len(initials) == 0 {
		log.Warnf("invalid syllabary: %s", syllabary)
		return nil, domain.WithInvalidArgument("invalid syllabary: Consonant Not Request")
	}

	cards, err := u.cardRepository.ListCardsByInitial(ctx, initials)
	if err != nil {
		log.Errorf("failed to get cards by syllabary: %v", err)
	}
	return cards, err
}
