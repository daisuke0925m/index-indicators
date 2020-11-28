package controllers

import (
	"fmt"
	"index-indicator-apis/app/models"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	fmt.Println(models.CreateNewFgis()) //TODO 定期実行(米株市場毎営業日の前後)

}
