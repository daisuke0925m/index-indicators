package controllers

import (
	"index-indicator-apis/server/app/models"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	go models.CreateNewFgis() //TODO 定期実行(米株市場毎営業日の前後)

}
