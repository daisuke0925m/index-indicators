package main

import (
	"index-indicator-apis/server/app/controllers"
)

func main() {
	// mysql.CheckIsDb()
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
