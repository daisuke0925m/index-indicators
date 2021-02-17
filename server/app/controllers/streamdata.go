package controllers

import (
	"fmt"
	"index-indicators/server/app/models"

	"github.com/robfig/cron/v3"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {
	c := cron.New()

	// 平日23:30 TODO米国平日の市場取引時間
	c.AddFunc("30 23 * * 1-5", func() {
		fmt.Println("handle fgis")
		models.CreateNewFgis()
		fmt.Println("handle ticers")
		models.SaveTickers()
		fmt.Println("finish")
	})
	c.Start()

}
