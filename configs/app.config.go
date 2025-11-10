package configs

import (
	"errors"
	"os"
)

type AppConfig struct {
	ServerPort            string
	DSN                   string
	Secret                string
	TwilioAccount         string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
}

func SetUpEnv() (cfg AppConfig, err error) {
	httpPort := os.Getenv("PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("no enviroment variables found")
	}

	dsn := os.Getenv("DSN")

	if len(dsn) < 1 {
		return AppConfig{}, errors.New("no enviroment variables found")
	}

	secret := os.Getenv("SECRET")

	if len(secret) < 1 {
		return AppConfig{}, errors.New("no enviroment variables found")
	}

	twilio_sid := os.Getenv("TWILIO_ACCOUNT_SID")
	twilio_auth_token := os.Getenv("TWILIO_AUTH_TOKEN")
	twilio_from_phone_number := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	if len(twilio_sid) < 1 || len(twilio_auth_token) < 1 || len(twilio_from_phone_number) < 1 {
		return AppConfig{}, errors.New("no enviroment variables found")
	}

	return AppConfig{ServerPort: httpPort, DSN: dsn, Secret: secret, TwilioAccount: twilio_sid, TwilioAuthToken: twilio_auth_token, TwilioFromPhoneNumber: twilio_from_phone_number}, nil

}
