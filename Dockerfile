### ビルドステージ ###
FROM golang:1.22 as builder

# 必要な環境変数
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# go.mod & go.sum を先にコピーして依存取得（キャッシュ効かせるため）
COPY go.mod go.sum ./
RUN go mod download

# 全体をコピー
COPY . .

# main.go の正確な位置に注意（cmd/main.go）

RUN go build -o main . \
  && ls -lh /app/main \
  && file /app/main

### ランタイムステージ ###
FROM alpine:3.18

COPY .env .

RUN apk update && \
    apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]