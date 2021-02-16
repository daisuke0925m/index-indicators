# index_indicators

```
git clone

config.ini .env を設置

docker compose up
```

## 起動コマンド

### docker
`docker-compose up`

### ローカルPC
環境変数はenvファイルまたはエディタのコンフィグで管理
ターミナルで実行する場合は下記コマンド
```
SRC_ROOT=$PWD/ \
MYSQL_HOST=localhost  \
REDIS_HOST=localhost  \
API_URL=http://localhost:3000  \
MYSQL_DATABASE=index_indicators  \
MYSQL_USER=index_indicators  \
MYSQL_PASSWORD=index_indicators  \
MYSQL_ROOT_PASSWORD=index_indicators  \
go run cmd/index-indicators/main.go
```

## テスト(ローカルPC)
`SRC_ROOT=$PWD/ go test -v ./テストしたいパッケージディレクトリ`
modelsパッケージの場合は
`SRC_ROOT=$PWD/ go test -v ./app/models`

## ECS

### login 
`aws ecr get-login-password --region region | docker login --username AWS --password-stdin ID`
### build
`docker build -t index_indicators:v1 .`

### tag
`docker tag index_indicators:v1 ID/index_indicators:latest`

### ECR push
`docker push ID/index_indicators:latest`

### task definition
`aws ecs register-task-definition --cli-input-json file://task-definition.json`

### create service
`aws ecs create-service --cli-input-json file://ecs-service.json`
