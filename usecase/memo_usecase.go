package usecase

import (
	"context"

	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"gorm.io/gorm"
)

type MemoUsecase interface {
	GetMemo(ctx context.Context, id uint64) (*model.Memo, error)
	ListMemos(ctx context.Context) ([]model.Memo, error)
	CreateMemo(ctx context.Context, memo *model.Memo) error
	UpdateMemo(ctx context.Context, memo *model.Memo) error
	DeleteMemo(ctx context.Context, id uint64) error
}

// usecase/memo_usecase.go
type memoUsecase struct {
	db             *gorm.DB
	memoRepository repository.MemoRepository
}

func NewMemoUsecase(db *gorm.DB, memoRepo repository.MemoRepository) MemoUsecase {
	return &memoUsecase{
		db:             db,
		memoRepository: memoRepo,
	}
}

func (u *memoUsecase) GetMemo(ctx context.Context, id uint64) (*model.Memo, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("GetMemo called with ID=%d", id)

	memo, err := u.memoRepository.GetMemo(ctx, id)
	if err != nil {
		log.Errorf("failed to get memo: %v", err)
	}
	return memo, err
}

func (u *memoUsecase) ListMemos(ctx context.Context) ([]model.Memo, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListMemos called")

	memos, err := u.memoRepository.ListMemos(ctx)
	if err != nil {
		log.Errorf("failed to list memos: %v", err)
	}
	return memos, err
}

func (u *memoUsecase) CreateMemo(ctx context.Context, memo *model.Memo) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("CreateMemo called: %+v", memo)

	err := u.memoRepository.CreateMemo(ctx, memo)
	if err != nil {
		log.Errorf("failed to create memo: %v", err)
	}
	return err
}

func (u *memoUsecase) UpdateMemo(ctx context.Context, memo *model.Memo) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("UpdateMemo called: ID=%d", memo.ID)

	err := u.memoRepository.UpdateMemo(ctx, memo)
	if err != nil {
		log.Errorf("failed to update memo: %v", err)
	}
	return err
}

func (u *memoUsecase) DeleteMemo(ctx context.Context, id uint64) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("DeleteMemo called: ID=%d", id)

	err := u.memoRepository.DeleteMemo(ctx, id)
	if err != nil {
		log.Errorf("failed to delete memo: %v", err)
	}
	return err
}
