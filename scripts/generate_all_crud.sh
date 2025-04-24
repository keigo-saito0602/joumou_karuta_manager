
# 権限の付与
# chmod +x scripts/generate_model.sh
# chmod +x scripts/generate_repository.sh
# chmod +x scripts/generate_usecase.sh
# chmod +x scripts/generate_handler.sh
# chmod +x scripts/generate_validator.sh
# chmod +x scripts/generate_all_crud.sh

# 📌 使い方:
# ./scripts/generate_all_crud.sh Memo
# → user という単語を元に CRUD 関連ファイルを生成

# 作成されるファイル一覧（例: User）
# domain/model/user.go
# domain/repository/user_repository.go
# infrastructure/repository/user_repository.go
# usecase/user_usecase.go
# interface/handler/user_handler.go

NAME=$1

./scripts/generate_model.sh $NAME
./scripts/generate_repository.sh $NAME
./scripts/generate_usecase.sh $NAME
./scripts/generate_handler.sh $NAME
./scripts/generate_validator.sh $NAME

echo "🔧 以下のファイルに $NAME の内容を追加してください。"
echo "di/container.go"
echo "di/handlers.go"
echo "di/repositories.go"
echo "di/usecases.go"
echo "router/route.go"
echo "interface/handler/handlers.go"

echo "必要な場合マイグレーションファイルを追加してください。"
echo "assets/migrations"
echo "例:"
echo "assets/migrations/202504210000000_create-$NAME.up.sql"
echo "assets/migrations/202504210000000_create-$NAME.down.sql"

echo "🎉 $NAME CRUD 全層生成完了！"
