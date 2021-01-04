package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// ConfigList 環境変数などを設定
type ConfigList struct {
	FgiAPIKey  string
	FgiAPIHost string
	Port       int
	JwtAccess  string
	JwtRefresh string
}

// Config グローバル定義
var Config ConfigList

func init() {
	cfg, err := ini.Load(os.Getenv("SRC_ROOT") + "config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		FgiAPIKey:  cfg.Section("fgi").Key("api_key").String(),
		FgiAPIHost: cfg.Section("fgi").Key("api_host").String(),
		Port:       cfg.Section("web").Key("port").MustInt(),
		JwtAccess:  cfg.Section("secret").Key("jwt_access").String(),
		JwtRefresh: cfg.Section("secret").Key("jwt_refresh").String(),
	}
}
