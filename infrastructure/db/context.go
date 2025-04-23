package db

import (
	"context"

	"gorm.io/gorm"
)

// dbContextKey gorm.DB 参照のコンテキスト入出力時のキー構造体（パッケージプライベートで秘匿）
type dbContextKey struct{}

// dbKey キー（パッケージプライベートで秘匿）
var dbKey = dbContextKey{}

// ToContext 指定された gorm.DB の参照を設定したコンテキストを返します。
func ToContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, tx)
}

// FromContext コンテキストから gorm.DB の参照を取得します。見つからない場合は nil を返ります。
func FromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(dbKey).(*gorm.DB)
	if ok {
		return tx
	}
	return nil
}
