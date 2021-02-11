# index_indicators

```
git clone

config.iniを設置

docker compose up
```

## 起動コマンド

### docker
`docker-compose up`

### ローカルPC
`SRC_ROOT=$PWD/ go run cmd/index_indicators/main.go`

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
