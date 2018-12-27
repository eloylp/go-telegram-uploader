package watcher

import (
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

var once sync.Once
var watcher *fsnotify.Watcher

func GetGenericWatcher(initialPath string) *fsnotify.Watcher {
	once.Do(func() {
		createdWatcher, err := fsnotify.NewWatcher()
		fails.FailIfError(err)
		err = createdWatcher.Add(initialPath)
		fails.FailIfError(err)
		watcher = createdWatcher
	})
	return watcher
}

func StartMediaWatcher(watcher *fsnotify.Watcher, handler func(string)) {
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
