package main

import (
	"fmt"

	"index-indicator-apis/config"
)

func main() {
	fmt.Println(config.Config.FgiAPIKey)
	fmt.Println(config.Config.FgiAPIHost)
}
