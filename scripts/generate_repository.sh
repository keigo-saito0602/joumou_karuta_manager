NAME=$1
LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

cat <<EOF > infrastructure/repository/${LOWER}_repository.go
package repository

import (
  "context"
  "log"

  "github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
  "gorm.io/gorm"
)

type ${NAME}Repository interface {
  Create${NAME}(ctx context.Context, m *model.${NAME}) error
  List${NAME}s(ctx context.Context) ([]model.${NAME}, error)
  Get${NAME}(ctx context.Context, id uint64) (*model.${NAME}, error)
  Update${NAME}(ctx context.Context, m *model.${NAME}) error
  Delete${NAME}(ctx context.Context, id uint64) error
}

type gorm${NAME}Repository struct {
  db *gorm.DB
}

func New${NAME}Repository(db *gorm.DB) ${NAME}Repository {
  return &gorm${NAME}Repository{db: db}
}

// 実装は省略
EOF

echo "✅ repository/${LOWER}_repository.go created"
