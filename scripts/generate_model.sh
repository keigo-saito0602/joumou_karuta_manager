


# 🏷️ GORMタグの主なオプション一覧
# 主キーとID関連
# primaryKey：主キーを指定します。​

# ID uint64 `gorm:"primaryKey"`
# autoIncrement：自動インクリメントを指定します。​

# ID uint64 `gorm:"autoIncrement"`
# タイムスタンプ管理
# autoCreateTime：レコード作成時に自動で現在時刻を設定します。​

# CreatedAt time.Time `gorm:"autoCreateTime"`
# autoUpdateTime：レコード更新時に自動で現在時刻を設定します。​
# GORM

# UpdatedAt time.Time `gorm:"autoUpdateTime"`
# autoCreateTime:nano / autoUpdateTime:milli：UNIXナノ秒やミリ秒でのタイムスタンプを指定します。

# カラム名と型の指定
# column：データベース上のカラム名を指定します。​

# Name string `gorm:"column:username"`
# type：データベース上のカラム型を指定します。​

# Age int `gorm:"type:int"`
# size：文字列型のサイズを指定します。​

# Name string `gorm:"size:255"`
# コメントと制約
# comment：カラムにコメントを追加します。​

# Name string `gorm:"comment:'ユーザー名'"`
# not null：NULLを許可しない制約を追加します。​

# Email string `gorm:"not null"`
# default：デフォルト値を指定します。​

# Status string `gorm:"default:'active'"`

# index：インデックスを作成します。​

# Email string `gorm:"index"`
# uniqueIndex：ユニークインデックスを作成します。​
# GORM

# Username string `gorm:"uniqueIndex"`
# index:,sort:desc：降順のインデックスを作成します。​
# GORM

# CreatedAt time.Time `gorm:"index:,sort:desc"`
# index:idx_name,priority:2：インデックス名と優先度を指定します。​
# GORM

# FirstName string `gorm:"index:idx_name,priority:2"`

NAME=$1
LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

cat <<EOF > domain/model/${LOWER}.go
package model

import "time"

type ${NAME} struct {
  ID        uint64    \`json:"id"\`
  Value1    uint64    \`json:"value1"\`
  Value2    *string    \`json:"value2"\`
  CreatedAt time.Time \`json:"created_at"\`
  UpdatedAt time.Time \`json:"updated_at"\`
}
EOF

echo "✅ model/${LOWER}.go created"