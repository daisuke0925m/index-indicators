package models

import (
	"fmt"
	"time"

	"github.com/markcheno/go-quote"
)

func dateToString(d time.Time) string {
	const layout = "2006-01-02"
	dString := d.Format(layout)
	return dString
}

// SaveTickers save tickers data
func SaveTickers() (err error) {
	tickers := []string{"spxl"}
	// tickers := []string{"spxl", "^skew", "tlt", "gld", "gldm", "spy"}
	today := time.Now()
	twoYAgo := today.AddDate(0, 0, -3)

	for i, t := range tickers {
		stockData, err := quote.NewQuoteFromYahoo(t, dateToString(twoYAgo), dateToString(today), quote.Daily, true)
		if err != nil {
			return err
		}
		fmt.Println(stockData)
		fmt.Println(i)
	}

	// fmt.Println(dateToString(today))
	// fmt.Println(dateToString(twoYAgo))
	return err
}
