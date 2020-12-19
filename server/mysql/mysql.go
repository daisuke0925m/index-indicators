package mysql

import (
	"index-indicator-apis/server/app/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func SQLConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "iia"
	PASS := "iia"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "index_indicator_apis"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}

//AutoMigrate マイグレーション
func AutoMigrate() {
	var err error
	db, err := SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&entity.Fgi{})
}
