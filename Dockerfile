### ビルドステージ ###
FROM --platform=linux/arm64 golang:1.22-alpine as builder

# 必要な環境変数
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

WORKDIR /app

# file コマンド追加のため Alpine パッケージを入れる
RUN apk update && apk add --no-cache file

# go.mod & go.sum を先にコピーして依存取得（キャッシュ効かせるため）
COPY go.mod go.sum ./
RUN go mod download

# 全体をコピー
COPY . .

# アプリケーションをビルドし、確認のためにサイズと種類を表示
RUN go build -o main . \
  && ls -lh /app/main \
  && file /app/main

### ランタイムステージ ###
FROM alpine:3.18

RUN apk update && \
    apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080