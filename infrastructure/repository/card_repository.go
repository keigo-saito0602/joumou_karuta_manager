package repository

import (
	"context"
	"log"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"gorm.io/gorm"
)

type CardRepository interface {
	ListCards(ctx context.Context) ([]model.Card, error)
	GetCard(ctx context.Context, id uint64) (*model.Card, error)
	ListCardsByInitial(ctx context.Context, initials []string) ([]model.Card, error)
}

type gormCardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &gormCardRepository{db: db}
}

func (r *gormCardRepository) ListCards(ctx context.Context) ([]model.Card, error) {
	var cards []model.Card
	err := r.db.Find(&cards).Error
	if err != nil {
		log.Printf("[CardRepository][ListCards] Failed to fetch cards: %v", err)
	}
	return cards, err
}

func (r *gormCardRepository) GetCard(ctx context.Context, id uint64) (*model.Card, error) {
	var card model.Card
	err := r.db.First(&card, id).Error
	if err != nil {
		log.Printf("[CardRepository][GetCard] Failed to find card with ID=%d: %v", id, err)
		return nil, err
	}
	return &card, nil
}

func (r *gormCardRepository) ListCardsByInitial(ctx context.Context, initials []string) ([]model.Card, error) {
	var cards []model.Card
	err := r.db.
		Where("initial IN ?", initials).
		Find(&cards).Error
	if err != nil {
		log.Printf("[CardRepository][ListCardsByInitials] Failed to find cards by initials: %v", err)
	}
	return cards, err
}
