package main

import (
	"index-indicator-apis/app/models"
	"index-indicator-apis/mysql"
)

func main() {
	mysql.CheckIsDb()
	models.CreateNewFgis()
}
