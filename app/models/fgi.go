package models

import (
	"index-indicator-apis/mysql"
	"time"

	"index-indicator-apis/config"
	"index-indicator-apis/fgi"
)

const (
	tableNameFgis = "fgis"
)

// Fgis 日足格納
type Fgis struct {
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
	db.AutoMigrate(&Fgis{})
}

// NewFgis fgi.StructFgiを受け取り、Fgisに変換して返す
func NewFgis(f fgi.StructFgi) *Fgis {
	createdAt := time.Now()
	nowValue := f.Fgi.Current.Value
	nowText := f.Fgi.Current.ValueText
	pcValue := f.Fgi.PreviousClose.Value
	pcText := f.Fgi.PreviousClose.ValueText
	oneWValue := f.Fgi.OneWeekAgo.Value
	oneWText := f.Fgi.OneWeekAgo.ValueText
	oneMValue := f.Fgi.OneMonthAgo.Value
	oneMText := f.Fgi.OneMonthAgo.ValueText
	oneYValue := f.Fgi.OneYearAgo.Value
	oneYText := f.Fgi.OneYearAgo.ValueText

	return &Fgis{
		createdAt,
		nowValue,
		nowText,
		pcValue,
		pcText,
		oneWValue,
		oneWText,
		oneMValue,
		oneMText,
		oneYValue,
		oneYText,
	}
}

// Create NewFgisをsaveする
func (f *Fgis) Create() error {
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(&f)
	return err
}

// CreateNewFgis migration後にapiを叩きdbに保存する
func CreateNewFgis() {
	// initFgis()　migration
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	f, _ := fgiClient.GetFgi()
	fgi := NewFgis(f)
	(fgi.Create())
}
