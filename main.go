package main

import (
	"github.com/eloylp/go-telegram-uploader/config"
	"github.com/eloylp/go-telegram-uploader/watcher"
)

func main() {
	watcher.Watcher(config.GetConfig().FolderToScan)
}
