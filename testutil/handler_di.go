package testutil

import (
	"github.com/golang/mock/gomock"
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler"
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler/mocks"
	"github.com/keigo-saito0602/joumou_karuta_manager/validation"
)

// TestHandlerDI provides dependencies for multiple handler tests
type TestHandlerDI struct {
	Ctrl          *gomock.Controller
	UserUsecase   *mocks.MockUserUsecase
	MemoUsecase   *mocks.MockMemoUsecase
	UserValidator *validation.UserValidator
	MemoValidator *validation.MemoValidator
	UserHandler   *handler.UserHandler
	MemoHandler   *handler.MemoHandler
}

// NewTestHandlerDI sets up mock dependencies for handler tests
func NewTestHandlerDI(t gomock.TestReporter) *TestHandlerDI {
	ctrl := gomock.NewController(t)

	userUsecase := mocks.NewMockUserUsecase(ctrl)
	memoUsecase := mocks.NewMockMemoUsecase(ctrl)

	userValidator := validation.NewUserValidator()
	memoValidator := validation.NewMemoValidator(userUsecase)

	userHandler := handler.NewUserHandler(userUsecase, userValidator)
	memoHandler := handler.NewMemoHandler(memoUsecase, memoValidator)

	return &TestHandlerDI{
		Ctrl:          ctrl,
		UserUsecase:   userUsecase,
		MemoUsecase:   memoUsecase,
		UserValidator: userValidator,
		MemoValidator: memoValidator,
		UserHandler:   userHandler,
		MemoHandler:   memoHandler,
	}
}

// Cleanup finishes the gomock controller
func (di *TestHandlerDI) Cleanup() {
	di.Ctrl.Finish()
}
