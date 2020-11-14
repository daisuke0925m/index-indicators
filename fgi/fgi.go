// package fgi

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// const baseUrl := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"
// func main() {

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("x-rapidapi-host", "")
// 	req.Header.Add("x-rapidapi-key", "")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println(res)
// 	fmt.Println(string(body))

// }