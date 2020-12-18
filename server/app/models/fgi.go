package models

import (
	"fmt"
	"index-indicator-apis/server/mysql"
	"time"

	"index-indicator-apis/server/config"
	"index-indicator-apis/server/fgi"
)

// Fgi 日足格納
type Fgi struct {
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

// NewFgis fgi.StructFgiを受け取り、Fgisに変換して返す
func NewFgis(f fgi.StructFgi) *Fgi {
	return &Fgi{
		time.Now(),
		f.Fgi.Current.Value,
		f.Fgi.Current.ValueText,
		f.Fgi.PreviousClose.Value,
		f.Fgi.PreviousClose.ValueText,
		f.Fgi.OneWeekAgo.Value,
		f.Fgi.OneWeekAgo.ValueText,
		f.Fgi.OneMonthAgo.Value,
		f.Fgi.OneMonthAgo.ValueText,
		f.Fgi.OneYearAgo.Value,
		f.Fgi.OneYearAgo.ValueText,
	}
}

// Create NewFgisをsaveする
func (f *Fgi) Create() error {
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(&f)
	return err
}

// CreateNewFgis migration後にapiを叩きdbに保存する
func CreateNewFgis() error {
	initFgis() //migration
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	f, err := fgiClient.GetFgi()
	if err != nil {
		return err
	}
	fgi := NewFgis(f)
	fmt.Println(fgi)
	fmt.Println(fgi.Create()) //TODO 日本時間の定期実行/contorollers/streamIngestionData
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
