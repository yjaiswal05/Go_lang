package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "<Your Slack bot token>")
	os.Setenv("CHANNEL_ID", "<Your channel token>")

	
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")
	fileArr := []string{"videoframe_4094532.png"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileV2Parameters{
			Channel: channel,
			File: fileArr[i],
			Filename :"videoframe_4094532.png",
			FileSize: 1,
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name %s",file)
	}
}