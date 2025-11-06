package configs

import (
	"errors"
	"os"
)

type AppConfig struct {
	ServerPort string
	DSN        string
	Secret     string
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

	return AppConfig{ServerPort: httpPort, DSN: dsn, Secret: secret}, nil

}
