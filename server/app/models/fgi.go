package models

import (
	"index-indicators/server/app/entity"
	"log"
	"os"
	"time"

	"index-indicators/server/fgi"
)

// CreateNewFgis migration後にapiを叩きdbに保存する
func (m *Models) CreateNewFgis() error {
	fgiClient := fgi.New(os.Getenv("FGI_KEY"), os.Getenv("FGI_HOST"))
	f, err := fgiClient.GetFgi()
	if err != nil {
		return err
	}
	log.Printf("insert value:%v\n", f.Fgi)

	m.DB.Create(&entity.Fgi{
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
func (m *Models) GetFgis(limit int) []entity.Fgi {

	fgis := []entity.Fgi{}
	m.DB.Order("created_at desc").Limit(limit).Find(&fgis)
	return fgis
}
