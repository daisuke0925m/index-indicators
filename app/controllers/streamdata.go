package controllers

import (
	"index-indicator-apis/app/models"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	models.CreateNewFgis() //TODO 定期実行

}
