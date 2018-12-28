# Go telegram uploader

Upload your media , videos or pictures syncing a folder in your system with your telegram bot.
It will detect file creation.

## How to build
There is one script at the root of the project, [build_all.sh](build_all.sh) that will build all binaries placing them 
in the [build folder](build)
```bash
./build_all.sh
```
Will be enough if you have the [go toolchain](https://golang.org/doc/install) installed on your system.

## How to run
```bash
./go-telegram-uploader --chat=your_chat_id --folder=/your/folder --token=your_bot_token 
```
## Precompiled binaries
If you trust on this repo, you can download latest binaries from the current or previous [releases](https://github.com/eloylp/go-telegram-uploader/releases)
