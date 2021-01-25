package controllers

import (
	"index-indicator-apis/server/app/models"

	"github.com/robfig/cron/v3"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	models.SaveTickers()
	c := cron.New()

	// 平日23:30 TODO米国平日の市場取引時間
	c.AddFunc("30 23 * * 1-5", func() {
		models.CreateNewFgis()
		models.SaveTickers()
	})
	c.Start()

}
