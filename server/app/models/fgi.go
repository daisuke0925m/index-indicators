package models

import (
	"fmt"
	"index-indicators/server/app/entity"
	"index-indicators/server/db"
	"os"
	"time"

	"index-indicators/server/fgi"
)

// CreateNewFgis migration後にapiを叩きdbに保存する
func CreateNewFgis() error {
	fgiClient := fgi.New(os.Getenv("FGI_KEY"), os.Getenv("FGI_HOST"))
	f, err := fgiClient.GetFgi()
	if err != nil {
		return err
	}
	fmt.Printf("insert value:%v\n", f.Fgi)

	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(&entity.Fgi{
		CreatedAt: time.Now(),
		NowValue:  f.Fgi.Current.Value,
		NowText:   f.Fgi.Current.ValueText,
		PcValue:   f.Fgi.PreviousClose.Value,
		PcText:    f.Fgi.PreviousClose.ValueText,
		OneWValue: f.Fgi.OneWeekAgo.Value,
		OneWText:  f.Fgi.OneWeekAgo.ValueText,
		OneMValue: f.Fgi.OneMonthAgo.Value,
		OneMText:  f.Fgi.OneMonthAgo.ValueText,
		OneYValue: f.Fgi.OneYearAgo.Value,
		OneYText:  f.Fgi.OneYearAgo.ValueText,
	})

	return err
}

// GetFgis api for webserver
func GetFgis(limit int) []entity.Fgi {
	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fgis := []entity.Fgi{}
	db.Order("created_at desc").Limit(limit).Find(&fgis)
	return fgis
}
