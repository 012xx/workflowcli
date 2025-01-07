package main

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	toUserID := os.Getenv("LINE_TO_USER_ID")
}

if channelSecret == "" || channelToken == "" || toUserID == "" {
	log.Fatal("環境変数 LINE_CHANNEL_SECRET, LINE_CHANNEL_TOKEN, LINE_TO_USER_ID を設定してください")
}