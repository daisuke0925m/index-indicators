FROM golang:1.15 as builder

# TODO GO env設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src/app/server

# TODO
# COPY wait-for-it.sh ./
# RUN chmod +x /usr/wait-for-it.sh

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/index-indicator-apis/main.go

FROM alpine

RUN apk add --no-cache ca-certificates && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

COPY --from=builder usr/src/app/server/main usr/src/app/server/config.ini ./

# CMD ["./main"]
# デバック用↓
CMD ["sleep", "5000"]
