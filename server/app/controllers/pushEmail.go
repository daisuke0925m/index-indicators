package controllers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/robfig/cron/v3"
)

var sender = "noreply@index-indicators.com"
var title = "昨日の株価・インディケーターをお知らせします。"
var mailBody = "登録済みの株価・インディケーター ⬇️\n"
var mailBodyFgi = ""
var mailBodyFooter = "通知銘柄・イディケーターの変更はこちら\nhttps://mt.index-indicators.com"

func initEmail(to string, title string, body string) error {
	awsSession := session.New(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_SES_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_SES_ACCESS_KEY"), os.Getenv("AWS_SES_SECRETE_KEY"), ""),
	})
	svc := ses.New(awsSession)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(title),
			},
		},
		Source: aws.String(sender),
	}
	_, err := svc.SendEmail(input)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// PushEmail pushing email for indicators notice
func (a *App) createEmail() error {
	users, err := a.DB.GetAllUsers()
	if err != nil {
		log.Println("cloud not get users, db error.")
	}

	for _, user := range users {
		to := user.Email

		likes, err := a.DB.FindUsersLikes(user)
		if err != nil {
			log.Printf("cloud not get %v 's likes, db error. \n", user)
			return err
		}

		if len(likes) > 0 {
			for _, like := range likes {
				if like.Symbol == "fgi" {
					limit := 1
					fgis := a.DB.GetFgis(limit)
					lastFgi := fgis[0]
					mailBodyFgi = "\n-------------\n" + "Fear & Greed Index\n" +
						strings.Split(lastFgi.CreatedAt.String(), " ")[0] +
						" " + lastFgi.NowText + " " + strconv.Itoa(lastFgi.NowValue) +
						"\n 1Week Ago" + lastFgi.NowText + " " + strconv.Itoa(lastFgi.NowValue) +
						"\n 1Month Ago" + lastFgi.OneWText + " " + strconv.Itoa(lastFgi.OneWValue) +
						"\n 1Year Ago" + lastFgi.OneWText + " " + strconv.Itoa(lastFgi.OneWValue) +
						"\n-------------\n"
				} else {
					tickers, err := a.DB.GetTickerAll(like.Symbol)
					if err != nil {
						log.Printf("cloud not get %v 's ticker data, db error. \n", like)
						return err
					}
					latestData := tickers[len(tickers)-1]
					body := "\n-------------\n" +
						latestData.Symbol +
						"\nopen " + strconv.FormatFloat(latestData.Open, 'f', 0, 64) +
						"\nhigh " + strconv.FormatFloat(latestData.High, 'f', 0, 64) +
						"\nLow " + strconv.FormatFloat(latestData.Low, 'f', 0, 64) +
						"\nClose " + strconv.FormatFloat(latestData.Close, 'f', 0, 64) +
						"\n-------------\n"

					mailBody = mailBody + body
					fmt.Printf("-------%v", mailBody)
				}
			}
			err = initEmail(to, title, mailBody+mailBodyFgi+mailBodyFooter)
			if err != nil {
				log.Printf("mail sending error to \n username=%v email=%v", user.UserName, user.Email)
			}
			mailBody = "登録済みの株価・インディケーター ⬇️\n"
		}

	}
	return nil
}

// PushEmail 毎朝プッシュメール通知
func (a *App) PushEmail() {
	err := a.createEmail()
	if err != nil {
		log.Print(err.Error())
	}
	c := cron.New()

	// 平日 AM9:00
	c.AddFunc("00 09 * * 1-5", func() {
		log.Println("pushing email")
	})
	c.Start()

}
