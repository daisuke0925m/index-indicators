package main

import (
	"index-indicator-apis/server/app/controllers"
	"index-indicator-apis/server/db"
)

func main() {
	db.AutoMigrate()
	db.InitRedis()
	go controllers.StreamIngestionData()
	controllers.StartWebServer()
}
