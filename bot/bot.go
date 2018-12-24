package bot

import (
	"github.com/eloylp/go-telegram-uploader/config"
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
)

var bot *tgbotapi.BotAPI
var once sync.Once

func GetBot() *tgbotapi.BotAPI {
	once.Do(func() {
		telegramToken := config.GetConfig().BotToken
		createdBot, err := tgbotapi.NewBotAPI(telegramToken)
		fails.FailIfError(err)
		bot = createdBot
	})
	return bot
}

func SendToBot(chattable tgbotapi.Chattable) {
	_, err := GetBot().Send(chattable)
	log.Println(err)
}
