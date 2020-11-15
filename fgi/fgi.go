package fgi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// FgiStruct fgi格納
type FgiStruct struct {
	Fgi struct {
		Now struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"now"`
		PreviousClose struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"previousClose"`
		OneWeekAgo struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"oneWeekAgo"`
		OneMonthAgo struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"oneMonthAgo"`
		OneYearAgo struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"oneYearAgo"`
	} `json:"fgi"`
}

// GetFgi api実行
func (fgi *FgiClient) GetFgi() {

	url := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", fgi.host)
	req.Header.Add("x-rapidapi-key", fgi.key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var fgiStruct FgiStruct
	if err := json.Unmarshal(body, &fgiStruct); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", fgiStruct.Fgi.Now.Value)
}
