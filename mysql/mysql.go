package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnection grobal
var DbConnection *sql.DB

func checkIsDb() {
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

	DbConnection, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/index_indicator_apis")
	if err != nil {
		panic(err)
	}
	defer DbConnection.Close()
	// cmd := `CREATE TABLE IF NOT EXISTS index_indicator_apis`
	// _, err := DbConnection.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	return
}

// func gormConnect() *gorm.DB {
// 	USER := "root"
// 	PASS := ""
// 	PROTOCOL := "tcp(127.0.0.1:3306)"
// 	DBNAME := "index_indicator_apis"

// 	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// ConnectMysql DB存在チェック後に接続
func ConnectMysql() {
	checkIsDb()
	// db := gormConnect()
	// defer db.Close()
}
