package main

import (
	"index-indicator-apis/server/app/controllers"
	"index-indicator-apis/server/mysql"
)

func main() {
	mysql.CheckIsDb()
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
