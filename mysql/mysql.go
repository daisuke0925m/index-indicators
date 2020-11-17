package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DbConnection grobal
var DbConnection *sql.DB

// CheckIsDb DBチェック
func CheckIsDb() {
	DbConnection, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
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

// SqlConnect　DB接続
func SqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "index_indicator_apis"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
