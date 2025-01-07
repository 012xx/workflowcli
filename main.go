package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env ファイルの読み込みに失敗しました")
	}

	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	toUserID := os.Getenv("LINE_CHANNEL_ID")

	if channelSecret == "" || channelToken == "" || toUserID == "" {
		log.Fatal("環境変数 LINE_CHANNEL_SECRET, LINE_CHANNEL_TOKEN, LINE_CHANNEL_ID を設定してください")
	}

	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go \" MTGが始まったよ！\"")
		return
	}

	messeage := os.Args[1]

	_, err = bot.PushMessage(toUserID, linebot.NewTextMessage(messeage)).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("メッセージを送信しました:", messeage)
}
