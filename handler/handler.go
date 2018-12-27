package handler

import (
	"errors"
	"github.com/eloylp/go-telegram-uploader/bot"
	"github.com/eloylp/go-telegram-uploader/config"
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func ProcessFile(path string) {

	pictureConfig, err := checkForPicture(path)
	if err == nil {
		bot.SendToBot(pictureConfig)
		return
	}
	videoConfig, err := checkForVideo(path)
	if err == nil {
		bot.SendToBot(videoConfig)
		return
	}
}

func checkForVideo(path string) (tgbotapi.VideoConfig, error) {
	fileName := filepath.Base(path)
	matched, err := regexp.MatchString(`(?i).*\.(avi|mp4|flv|mov)`, fileName)
	fails.FailIfError(err)
	if matched {
		data, err := ioutil.ReadFile(path)
		fails.FailIfError(err)
		photoData := tgbotapi.FileBytes{Name: fileName, Bytes: data}
		return tgbotapi.NewVideoUpload(config.GetConfig().ChatId, photoData), nil
	}
	return tgbotapi.VideoConfig{}, errors.New("cannot map for video")
}

func checkForPicture(path string) (tgbotapi.PhotoConfig, error) {
	fileName := filepath.Base(path)
	matched, err := regexp.MatchString(`(?i).*\.(jpeg|jpg|png)`, fileName)
	fails.FailIfError(err)
	if matched {
		data, err := ioutil.ReadFile(path)
		fails.FailIfError(err)
		photoData := tgbotapi.FileBytes{Name: fileName, Bytes: data}
		return tgbotapi.NewPhotoUpload(config.GetConfig().ChatId, photoData), nil
	}
	return tgbotapi.PhotoConfig{}, errors.New("cannot map for picture")
}
