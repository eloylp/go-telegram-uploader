package main

import (
	"github.com/eloylp/go-telegram-uploader/config"
	"github.com/eloylp/go-telegram-uploader/watcher"
)

func main() {
	done := make(chan bool)
	go watcher.Watcher(config.GetConfig().FolderToScan)
	<-done
}
