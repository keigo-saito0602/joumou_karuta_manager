FROM golang:1.22-bullseye

# タイムゾーンと非対話設定
ENV TZ=Asia/Tokyo
ENV DEBIAN_FRONTEND=noninteractive

# 作業ディレクトリ
WORKDIR /go/src/go-clean-api

# 先に go.mod / go.sum をコピーして依存解決
COPY go.mod go.sum ./
RUN go mod download

# 残りのファイルを全部コピー
COPY . .

# main.go のあるパッケージ（cmd）をビルド
RUN go build -o main ./cmd

# マイグレーションが必要なら、CMD引数で制御可能にしておくと柔軟
CMD ["./main"]
