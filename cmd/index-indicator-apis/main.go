package main

import (
	"index-indicator-apis/app/controllers"
)

func main() {
	// mysql.CheckIsDb()
	// controllers.StreamIngestionData()
	controllers.StartWebServer()
}
