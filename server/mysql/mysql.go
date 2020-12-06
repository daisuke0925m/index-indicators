package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// CheckIsDb DBチェック
func CheckIsDb() {
	DbConnection, err := sql.Open("mysql", "root:root@tcp(mysql_container:3306)/")
	if err != nil {
		panic(err)
	}
	defer DbConnection.Close()

	_, err = DbConnection.Exec("CREATE DATABASE IF NOT EXISTS index_indicator_apis")
	if err != nil {
		panic(err)
	}
	DbConnection.Close()

	return
}

// DbConnection grobal
var DbConnection *gorm.DB

// SQLConnect DB接続
func SQLConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql_container:3306)"
	DBNAME := "index_indicator_apis"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}
