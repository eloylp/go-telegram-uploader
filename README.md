# GO TELEGRAM UPLOADER

Upload your media , videos or pictures syncing a folder in your system with your telegram bot.
It will detect file creation.

## How to build
There is one script at the root of the project, [build_all.sh](build_all.sh) that will build all binaries placeing them 
at [build folder](build)
```bash
./build_all.sh
```
Will be enough if you have the [go toolchain](https://golang.org/doc/install) installed on your system.

## How to boot
```bash
BOT_TOKEN=your_token CHAT_ID=666 FOLDER_TO_SCAN=/path/folder ./go-telegram-uploader  
```
## Precompiled binaries
If you trust on this repo, you can download latest binaries from the current or [previous releases](https://github.com/eloylp/go-telegram-uploader/releases)