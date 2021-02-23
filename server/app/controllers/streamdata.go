package controllers

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// StreamIngestionData api保存を定期実行
func (a *App) StreamIngestionData() {
	c := cron.New()

	// 平日23:30 TODO米国平日の市場取引時間
	c.AddFunc("30 23 * * 1-5", func() {
		fmt.Println("handle fgis")
		a.DB.CreateNewFgis()
		fmt.Println("handle ticers")
		err := a.DB.SaveTickers()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("finish")
	})
	c.Start()

}
