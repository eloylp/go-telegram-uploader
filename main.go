package main

import (
	"github.com/eloylp/go-telegram-uploader/config"
	"github.com/eloylp/go-telegram-uploader/handler"
	"github.com/eloylp/go-telegram-uploader/watcher"
)

func main() {

	createdWatcher := watcher.Watcher(config.GetConfig().FolderToScan)
	defer createdWatcher.Close()
	done := make(chan bool)
	go watcher.StartMediaWatcher(createdWatcher, handler.ProcessFile)
	<-done
}
