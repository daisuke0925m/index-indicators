package fgi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"index-indicator-apis/config"
)

func DoRequest() {

	url := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", config.Config.FgiAPIHost)
	req.Header.Add("x-rapidapi-key", config.Config.FgiAPIKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
