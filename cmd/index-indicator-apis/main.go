package main

import (
	"fmt"
	"index-indicator-apis/config"
	"index-indicator-apis/fgi"
)

func main() {
	fgiClient := fgi.New(config.Config.FgiAPIKey, config.Config.FgiAPIHost)
	fmt.Println(fgiClient.GetFgi())
}
