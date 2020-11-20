package main

import (
	"index-indicator-apis/app/controllers"
	"index-indicator-apis/mysql"
)

func main() {
	mysql.CheckIsDb()
	controllers.StreamIngestionData()
}
