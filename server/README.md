# index-indicators

```
git clone

config.iniを設置

docker compose up
```

## 起動コマンド

### docker
`docker-compose up`

### ローカルPC
`SRC_ROOT=$PWD/ go run cmd/index-indicators/main.go`

## テスト(ローカルPC)
`SRC_ROOT=$PWD/ go test -v ./テストしたいパッケージディレクトリ`
modelsパッケージの場合は
`SRC_ROOT=$PWD/ go test -v ./app/models`

## ECS
### ECR push
`docker push 823425155155.dkr.ecr.ap-northeast-1.amazonaws.com/index_indicators:latest`

### task definition
`aws ecs register-task-definition --cli-input-json file://task-definition.json`

### create service
`aws ecs create-service --cli-input-json file://ecs-service.json`