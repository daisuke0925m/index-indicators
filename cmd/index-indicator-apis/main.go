package main

import (
	"index-indicator-apis/config"
	"index-indicator-apis/fgi"
)

func main() {
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	fgiClient.DoRequest()
}
