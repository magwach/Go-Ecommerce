package notification

import (
	"encoding/json"
	"errors"
	"go-ecommerce-app/configs"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phone, message string) error
}

type notificationClient struct {
	cfg configs.AppConfig
}

func (r *notificationClient) SendSMS(phone, message string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: r.cfg.TwilioAccount,
		Password: r.cfg.TwilioAuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+254736907046")
	params.SetFrom(r.cfg.TwilioFromPhoneNumber)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		return errors.New(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		log.Println(string(response))
	}
	return nil
}

func InitializeNotification(cfg configs.AppConfig) NotificationClient {
	return &notificationClient{
		cfg: cfg,
	}
}
