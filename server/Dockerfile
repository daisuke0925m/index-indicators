FROM golang:1.15 as builder

# TODO GO env設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src/app/server

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/index-indicator-apis/main.go

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder usr/src/app/server/main usr/src/app/server/config.ini ./

CMD ["./main"]
