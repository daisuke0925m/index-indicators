package controllers

import (
	"index-indicators/server/app/models"
	"log"

	"github.com/robfig/cron/v3"
)

// StreamIngestionData api保存を定期実行
func StreamIngestionData() {

	c := cron.New()

	// 平日23:30 TODO米国平日の市場取引時間
	c.AddFunc("00 00 * * *", func() {
		models, err := models.NewModels()
		if err != nil {
			log.Print(err.Error())
		}
		defer models.DB.Close()
		a := NewApp(models)

		log.Println("handle fgis")
		a.DB.CreateNewFgis()
		log.Println("handle ticers")
		err = a.DB.SaveTickers()
		if err != nil {
			log.Println("error")
		}
		log.Println("finish")
	})
	c.Start()

}
