package watcher

import (
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/eloylp/go-telegram-uploader/handler"
	"github.com/fsnotify/fsnotify"
	"log"
)

func Watcher(initialPath string) {
	createdWatcher, err := fsnotify.NewWatcher()
	fails.FailIfError(err)
	err = createdWatcher.Add(initialPath)
	fails.FailIfError(err)
	startWatcher(createdWatcher, handler.ProcessFile)
}

func startWatcher(watcher *fsnotify.Watcher, handler func(string)) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				filePath := event.Name
				go handler(filePath)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
