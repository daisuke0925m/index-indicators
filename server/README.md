# index-indicator-apis

## 起動コマンド

### docker 
`docker-compose up go`

### ローカルPC
`SRC_ROOT=$PWD/ go run cmd/index-indicator-apis/main.go`


## テスト(ローカルPC)
`SRC_ROOT=$PWD/ go test -v ./テストしたいパッケージディレクトリ`

modelsパッケージの場合は
`SRC_ROOT=$PWD/ go test -v ./app/models`

