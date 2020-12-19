package mysql

import (
	"fmt"
	"index-indicator-apis/server/app/entity"

	// Register for gorm
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

//AutoMigrate マイグレーション
func AutoMigrate() {
	fmt.Println("migrating database...")
	db, err := SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&entity.Fgi{})
	fmt.Println("finish migrate!")
}
