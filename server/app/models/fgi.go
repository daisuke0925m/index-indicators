package models

import (
	"index-indicator-apis/server/mysql"
	"time"

	"index-indicator-apis/server/config"
	"index-indicator-apis/server/fgi"
)

// Fgi 日足格納
type Fgi struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	NowValue  int       `json:"now_value,omitempty"`
	NowText   string    `json:"now_text,omitempty"`
	PcValue   int       `json:"pc_value,omitempty"`
	PcText    string    `json:"pc_text,omitempty"`
	OneWValue int       `json:"one_w_value,omitempty"`
	OneWText  string    `json:"one_w_text,omitempty"`
	OneMValue int       `json:"one_m_value,omitempty"`
	OneMText  string    `json:"one_m_text,omitempty"`
	OneYValue int       `json:"one_y_value,omitempty"`
	OneYText  string    `json:"one_y_text,omitempty"`
}

//initFgis マイグレーション
func initFgis() {
	var err error
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Fgi{})
}

// CreateNewFgis migration後にapiを叩きdbに保存する
func CreateNewFgis() error {
	initFgis() //migration
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	f, err := fgiClient.GetFgi()
	if err != nil {
		return err
	}

	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(&Fgi{
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
func GetFgis(limit int) []Fgi {
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fgis := []Fgi{}
	db.Order("created_at desc").Limit(limit).Find(&fgis)
	return fgis
}
