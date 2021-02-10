# TODO マルチステージビルド(timezone) wait-for-it up時のコンテナ停止を防ぐ(DBの接続制御)
FROM golang:1.15-alpine

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src/app/server

# COPY wait-for-it.sh ./
# RUN chmod +x /usr/wait-for-it.sh

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# RUN apt-get update && apt-get install vim -y

# FROM alpine

RUN apk add --no-cache ca-certificates && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# COPY --from=builder usr/src/app/server/main usr/src/app/server/config.ini ./

RUN go build -o main ./cmd/index-indicators/main.go

CMD ["/usr/src/app/server/./main"]
# デバック用↓
# CMD ["sleep", "5000"]

EXPOSE 8080
