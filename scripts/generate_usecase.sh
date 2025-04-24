NAME=$1
LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

cat <<EOF > usecase/${LOWER}_usecase.go
package usecase

import (
  "context"

  "github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
  dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
  "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
  "gorm.io/gorm"
)

type ${NAME}Usecase interface {
  Get${NAME}(ctx context.Context, id uint64) (*model.${NAME}, error)
  List${NAME}s(ctx context.Context) ([]model.${NAME}, error)
  Create${NAME}(ctx context.Context, m *model.${NAME}) error
  Update${NAME}(ctx context.Context, m *model.${NAME}) error
  Delete${NAME}(ctx context.Context, id uint64) error
}

type ${LOWER}Usecase struct {
  db *gorm.DB
  repo repository.${NAME}Repository
}

func New${NAME}Usecase(db *gorm.DB, repo repository.${NAME}Repository) ${NAME}Usecase {
  return &${LOWER}Usecase{db: db, repo: repo}
}

// 実装は省略
EOF

echo "✅ usecase/${LOWER}_usecase.go created"
