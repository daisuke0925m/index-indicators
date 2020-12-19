package main

import (
	"index-indicator-apis/server/app/controllers"
)

func main() {
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
