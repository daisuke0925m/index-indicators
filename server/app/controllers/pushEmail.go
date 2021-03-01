package controllers

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var sender = "noreply@index-indicators.com"

// PushEmail pushing email for indicators notice
func (a *App) PushEmail() {
	to := "送信先"
	title := "昨日の株価・インディケーター"
	body := "メール本文"
	err := initEmail(to, title, body)
	if err != nil {
		log.Println("mail sending error")
	}
}

// func (a *App) setEmails() error {
// }

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
