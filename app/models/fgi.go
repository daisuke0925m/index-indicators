package models

import (
	"index-indicator-apis/mysql"
)

const (
	tableNameFgis = "fgis"
)

//Init マイグレーション
func Init() {
	var err error
	db, err := mysql.SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Fgis{})
}

// Fgis 日足格納
type Fgis struct {
	ID        int    `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	NowValue  int    `json:"now_value,omitempty"`
	NowText   string `json:"now_text,omitempty"`
	PcValue   int    `json:"pc_value,omitempty"`
	PcText    string `json:"pc_text,omitempty"`
	OneWValue int    `json:"one_w_value,omitempty"`
	OneWText  string `json:"one_w_text,omitempty"`
	OneMValue int    `json:"one_m_value,omitempty"`
	OneMText  string `json:"one_m_text,omitempty"`
	OneYValue int    `json:"one_y_value,omitempty"`
	OneYText  string `json:"one_y_text,omitempty"`
}
