package models

import (
	"index-indicators/server/app/entity"
	"log"
	"time"

	"github.com/markcheno/go-quote"
)

func dateToString(d time.Time) string {
	const layout = "2006-01-02"
	dString := d.Format(layout)
	return dString
}

func (m *Models) createTickerRow(symbol string, date time.Time, open, high, low, close, volume float64) error {
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
	if err := m.DB.Create(&tickerRow).Error; err != nil {
		return err
	}

	return nil
}

func (m *Models) checkSymbolExist(symbol string) (bool, error) {
	var ticker entity.Ticker
	if err := m.DB.Where("symbol = ?", symbol).First(&ticker).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func (m *Models) checkLatestRecord(symbol string, date time.Time) (bool, error) {
	var ticker entity.Ticker
	if err := m.DB.Where("symbol = ? AND date = ?", symbol, date).First(&ticker).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func (m *Models) deleteLastRecord(symbol string, date time.Time) error {
	var ticker entity.Ticker
	if err := m.DB.Where("symbol = ? AND date = ?", symbol, date).Delete(&ticker).Error; err != nil {
		return err
	}

	return nil
}

// SaveTickers save tickers data
func (m *Models) SaveTickers() (err error) {
	symbols := []string{"spxl", "^skew", "tlt", "gld", "gldm", "spy"}
	today := time.Now()
	twoYAgo := today.AddDate(-2, 0, 0)

	for _, symbol := range symbols {
		// save 2years data
		tickerData, err := quote.NewQuoteFromYahoo(symbol, dateToString(twoYAgo), dateToString(today), quote.Daily, true)
		if err != nil {
			return err
		}
		// 監視用
		log.Print(tickerData.CSV() + "\n--" + symbol + "--\n")

		dataLength := len(tickerData.Open)
		len := dataLength - 1

		flag, err := m.checkSymbolExist(symbol)
		if err != nil {
			return err
		}

		// elseの場合のみ全ての結果をinsertする
		if flag {
			lastRecordDate := tickerData.Date[len]

			checkFlag, err := m.checkLatestRecord(symbol, lastRecordDate)
			if err != nil {
				return err
			}
			// dbにある最新のdateと取得したdateが一致した場合は処理を抜ける
			if !checkFlag {
				// 一致しなかった場合は一番古い日付のレコードを削除し最新のレコードをinsert
				err := m.deleteLastRecord(symbol, tickerData.Date[0])
				if err != nil {
					return err
				}
				err = m.createTickerRow(symbol, tickerData.Date[len], tickerData.Open[len], tickerData.High[len], tickerData.Low[len], tickerData.Close[len], tickerData.Volume[len])
				if err != nil {
					return err
				}
			}

		} else {

			for i := 0; i < dataLength; i++ {
				err := m.createTickerRow(symbol, tickerData.Date[i], tickerData.Open[i], tickerData.High[i], tickerData.Low[i], tickerData.Close[i], tickerData.Volume[i])
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

// GetTickerAll get 2years ticker data
func (m *Models) GetTickerAll(symbol string) ([]entity.Ticker, error) {
	var tickers []entity.Ticker

	if err := m.DB.Where("symbol = ?", symbol).Find(&tickers).Error; err != nil {
		return tickers, err
	}

	return tickers, nil
}
