package fgi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FgiClient api情報格納
type FgiClient struct {
	key        string
	host       string
	httpClient *http.Client
}

// New struct生成
func New(key, host string) *FgiClient {
	fgiClient := &FgiClient{key, host, &http.Client{}}
	return fgiClient
}

func (fgi *FgiClient) DoRequest() {

	url := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", fgi.host)
	req.Header.Add("x-rapidapi-key", fgi.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
