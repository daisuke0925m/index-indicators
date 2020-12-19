package controllers

import (
	"index-indicator-apis/server/app/models"
	"time"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			models.CreateNewFgis()
		}
	}

}
