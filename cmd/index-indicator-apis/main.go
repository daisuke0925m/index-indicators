package main

import (
	"index-indicator-apis/config"
	"index-indicator-apis/fgi"
	"index-indicator-apis/mysql"
)

func main() {
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	mysql.ConnectMysql()
	fgiClient.GetFgi()
}
