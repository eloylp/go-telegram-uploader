package watcher

import (
	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/fsnotify/fsnotify"
	"sync"
)

var once sync.Once
var watcher *fsnotify.Watcher

func GetWatcher() *fsnotify.Watcher {
	once.Do(func() {
		createdWatcher, err := fsnotify.NewWatcher()
		fails.FailIfError(err)
		watcher = createdWatcher
	})
	return watcher
}
