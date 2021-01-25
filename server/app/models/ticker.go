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

func createTickerRow(symbol string, date time.Time, open, high, low, close, volume float64) error {
	db, err := db.SQLConnect()
	if err != nil {
		return err
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

func checkSymbolExist(symbol string) (bool, error) {
	db, err := db.SQLConnect()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var ticker entity.Ticker
	if err := db.Where("symbol = ?", symbol).First(&ticker).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func checkLatestRecord(symbol string, date time.Time) (bool, error) {
	db, err := db.SQLConnect()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var ticker entity.Ticker
	if err := db.Where("symbol = ? AND date = ?", symbol, date).First(&ticker).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func deleteLastRecord(symbol string, date time.Time) error {
	db, err := db.SQLConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	var ticker entity.Ticker
	if err := db.Where("symbol = ? AND date = ?", symbol, date).Delete(&ticker).Error; err != nil {
		return err
	}

	return nil
}

// SaveTickers save tickers data
func SaveTickers() (err error) {
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
		len := dataLength - 1

		flag, err := checkSymbolExist(symbol)
		if err != nil {
			return err
		}

		// elseの場合のみ全ての結果をinsertする
		if flag {
			lastRecordDate := tickerData.Date[len]

			checkFlag, err := checkLatestRecord(symbol, lastRecordDate)
			if err != nil {
				return err
			}
			// dbにある最新のdateと取得したdateが一致した場合は処理を抜ける
			if !checkFlag {
				// 一致しなかった場合は一番古い日付のレコードを削除し最新のレコードをinsert
				err := deleteLastRecord(symbol, tickerData.Date[0])
				if err != nil {
					return err
				}
				err = createTickerRow(symbol, tickerData.Date[len], tickerData.Open[len], tickerData.High[len], tickerData.Low[len], tickerData.Close[len], tickerData.Volume[len])
				if err != nil {
					return err
				}
			}

		} else {

			for i := 0; i < dataLength; i++ {
				err := createTickerRow(symbol, tickerData.Date[i], tickerData.Open[i], tickerData.High[i], tickerData.Low[i], tickerData.Close[i], tickerData.Volume[i])
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

// GetTickerAll get 2years ticker data
func GetTickerAll(symbol string) ([]entity.Ticker, error) {
	var tickers []entity.Ticker
	db, err := db.SQLConnect()
	if err != nil {
		return tickers, err
	}
	defer db.Close()

	if err := db.Where("symbol = ?", symbol).Find(&tickers).Error; err != nil {
		return tickers, err
	}

	return tickers, nil
}
