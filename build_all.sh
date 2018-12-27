#!/usr/bin/env bash

HERE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
$HERE/go_multi_arch_build.sh github.com/eloylp/go-telegram-uploader $HERE/build 'windows/amd64|linux/amd64|linux/arm/5'
