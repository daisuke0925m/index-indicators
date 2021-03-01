package controllers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var sender = "noreply@index-indicators.com"
var title = "昨日の株価・インディケーターをお知らせします。"
var mailBody = "登録済みの株価・インディケーター ⬇️\n"
var mailBodyFooter = "通知銘柄・イディケーターの変更はこちら\nhttps://mt.index-indicators.com/"

// PushEmail pushing email for indicators notice
func (a *App) PushEmail() {
	users, err := a.DB.GetAllUsers()
	if err != nil {
		log.Println("cloud not get users, db error.")
	}
	fmt.Println(users)
	for _, user := range users {
		to := user.Email

		likes, err := a.DB.FindUsersLikes(user)
		if err != nil {
			log.Printf("cloud not get %v 's likes, db error. \n", user)
			return
		}

		if len(likes) > 0 {
			for _, like := range likes {
				if like.Symbol != "fgi" {
					tickers, err := a.DB.GetTickerAll(like.Symbol)
					if err != nil {
						log.Printf("cloud not get %v 's ticker data, db error. \n", like)
						return
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
				}
			}
		}

		fmt.Printf("-------%v", mailBody)
		err = initEmail(to, title, mailBody+mailBodyFooter)
		if err != nil {
			log.Printf("mail sending error to \n username=%v email=%v", user.UserName, user.Email)
		}
	}
}

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
