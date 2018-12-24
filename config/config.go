package config

import (
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/joeshaw/envdecode"
	"sync"
)

var config *Config
var once sync.Once

type Config struct {
	BotToken     string `env:"BOT_TOKEN,required"`
	ChatId       int64  `env:"CHAT_ID,required"`
	FolderToScan string `env:"FOLDER_TO_SCAN,required"`
}

func init() {
	config = &Config{}
}

func GetConfig() *Config {
	once.Do(func() {
		err := envdecode.Decode(config)
		fails.FailIfError(err)
	})
	return config
}
