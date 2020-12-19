package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func SQLConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "iia"
	PASS := "iia"
	PROTOCOL := "tcp(mysql_container:3306)"
	DBNAME := "index_indicator_apis"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}
