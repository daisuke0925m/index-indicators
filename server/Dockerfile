FROM golang:1.15 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /usr/src/app/server

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make

# -----------------------------------
FROM alpine

RUN apk add --no-cache ca-certificates && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

COPY --from=builder usr/src/app/server/main usr/src/app/server/config.ini ./

CMD ["./main"]


EXPOSE 8080
