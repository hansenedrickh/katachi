package config

import (
	"github.com/hansenedrickh/keisatsu"

	"github.com/hansenedrickh/katachi/utils"
)

type KeisatsuConfig struct {
	AppName     string
	WebhookURL  string
	SecretToken string
}

func newKeisatsuConfig() *KeisatsuConfig {
	return &KeisatsuConfig{
		AppName:     utils.FatalGetString("KEISATSU_APP_NAME"),
		WebhookURL:  utils.FatalGetString("KEISATSU_WEBHOOK_URL"),
		SecretToken: utils.FatalGetString("KEISATSU_SECRET_TOKEN"),
	}
}

func (k KeisatsuConfig) InitKeisatsu() keisatsu.Service {
	return keisatsu.New(k.AppName, k.WebhookURL, k.SecretToken)
}
