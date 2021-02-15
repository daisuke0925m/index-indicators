package db

import (
	"os"

	// Register for gorm
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func SQLConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "index_indicators"
	PASS := "index_indicators"
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":3306)"
	DBNAME := "index_indicators"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}
