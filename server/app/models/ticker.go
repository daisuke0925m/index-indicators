package models

import (
	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/db"
	"time"

	"github.com/markcheno/go-quote"
)

func dateToString(d time.Time) string {
	const layout = "2006-01-02"
	dString := d.Format(layout)
	return dString
}

func createTickerRow(symbol string, date time.Time, open, high, low, close, volume float64) (err error) {
	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	tickerRow := &entity.Ticker{
		Symbol:    symbol,
		Date:      date,
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		CreatedAt: time.Now(),
	}

	if err := db.Create(&tickerRow).Error; err != nil {
		return err
	}
	return nil
}

// SaveTickers save tickers data
func SaveTickers() (err error) {
	// symbols := []string{"spxl"}
	symbols := []string{"spxl", "^skew", "tlt", "gld", "gldm", "spy"}
	today := time.Now()
	twoYAgo := today.AddDate(-2, 0, 0)

	for _, symbol := range symbols {
		// save 2years data
		tickerData, err := quote.NewQuoteFromYahoo(symbol, dateToString(twoYAgo), dateToString(today), quote.Daily, true)
		if err != nil {
			return err
		}

		dataLength := len(tickerData.Open)
		for i := 0; i < dataLength; i++ {
			createTickerRow(symbol, tickerData.Date[i], tickerData.Open[i], tickerData.High[i], tickerData.Low[i], tickerData.Close[i], tickerData.Volume[i])
		}
	}

	// fmt.Println(dateToString(today))
	// fmt.Println(dateToString(twoYAgo))
	return err
}
