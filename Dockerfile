FROM golang:1.19-bullseye
ENV TZ=Asia/Jakarta
ENV DEBIAN_FRONTEND noninteractive

COPY . /app/
WORKDIR /app/

RUN go build app/main.go

RUN ./main --migrate

CMD ["./main"]
