package fgi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// APIClientFgi api情報格納
type APIClientFgi struct {
	key        string
	host       string
	httpClient *http.Client
}

// New struct生成
func New(key, host string) *APIClientFgi {
	fgiClient := &APIClientFgi{key, host, &http.Client{}}
	return fgiClient
}

// StructFgi fgi格納
type StructFgi struct {
	Fgi struct {
		Current struct {
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
func (fgi *APIClientFgi) GetFgi() (StructFgi, error) {

	url := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", fgi.host)
	req.Header.Add("x-rapidapi-key", fgi.key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return StructFgi{}, nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return StructFgi{}, nil
	}

	var fgiStruct StructFgi
	if err := json.Unmarshal(body, &fgiStruct); err != nil {
		log.Fatal(err)
	}

	return fgiStruct, nil
}
