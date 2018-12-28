package config

import (
	"flag"
	"os"
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

		flag.StringVar(&config.BotToken, "token", "", "The telegram bot token.")
		flag.Int64Var(&config.ChatId, "chat", 0, "The chat id to dump content")
		flag.StringVar(&config.FolderToScan, "folder", "", "The folder to keep scanning for media files.")
		flag.Parse()

		if config.BotToken == "" || config.ChatId == 0 || config.FolderToScan == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
	})
	return config
}
