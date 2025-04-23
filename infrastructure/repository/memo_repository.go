package repository

import (
	"context"
	"log"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"gorm.io/gorm"
)

type MemoRepository interface {
	CreateMemo(ctx context.Context, memo *model.Memo) error
	ListMemos(ctx context.Context) ([]model.Memo, error)
	GetMemo(ctx context.Context, id uint64) (*model.Memo, error)
	UpdateMemo(ctx context.Context, memo *model.Memo) error
	DeleteMemo(ctx context.Context, id uint64) error
}

type gormMemoRepository struct {
	db *gorm.DB
}

func NewMemoRepository(db *gorm.DB) MemoRepository {
	return &gormMemoRepository{db: db}
}

func (r *gormMemoRepository) CreateMemo(ctx context.Context, memo *model.Memo) error {

	err := r.db.Create(memo).Error
	if err != nil {
		log.Printf("[MemoRepository][CreateMemo] Failed to create memo: %v", err)
	}
	return err
}

func (r *gormMemoRepository) ListMemos(ctx context.Context) ([]model.Memo, error) {
	var memos []model.Memo
	err := r.db.Find(&memos).Error
	if err != nil {
		log.Printf("[MemoRepository][ListMemos] Failed to fetch memos: %v", err)
	}
	return memos, err
}

func (r *gormMemoRepository) GetMemo(ctx context.Context, id uint64) (*model.Memo, error) {
	var memo model.Memo
	err := r.db.First(&memo, id).Error
	if err != nil {
		log.Printf("[MemoRepository][GetMemo] Failed to find memo with ID=%d: %v", id, err)
		return nil, err
	}
	return &memo, nil
}

func (r *gormMemoRepository) UpdateMemo(ctx context.Context, memo *model.Memo) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Memo{}).
			Where("id = ?", memo.ID).
			Updates(map[string]interface{}{
				"user_id": memo.UserID,
				"title":   memo.Title,
				"content": memo.Content,
				"active":  memo.Active,
			}).Error; err != nil {
			log.Printf("[MemoRepository][UpdateMemo] Failed to update memo ID=%d: %v", memo.ID, err)
			return err
		}

		if err := tx.First(memo, memo.ID).Error; err != nil {
			log.Printf("[MemoRepository][UpdateMemo] Failed to refetch updated memo ID=%d: %v", memo.ID, err)
			return err
		}

		return nil
	})
}

func (r *gormMemoRepository) DeleteMemo(ctx context.Context, id uint64) error {
	err := r.db.Delete(&model.Memo{}, id).Error
	if err != nil {
		log.Printf("[MemoRepository][DeleteMemo] Failed to delete memo ID=%d: %v", id, err)
	}
	return err
}
