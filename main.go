package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {

	fmt.Println("Slackbot for File Uploading starting...")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5874689618118-6454218064391-y4yp55lKAb3KnKMpWXnf6hAy")
	os.Setenv("CHANNEL_ID", "C05RX9Q0V26")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr:=[]string{os.Getenv("CHANNEL_ID")}
	fileArr:=[]string{"file.txt"}
	 for i:=0 ; i<len(fileArr); i++ {
		params:=slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		uploadedFile, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf(" line 24 %s\n",err)
		}
		fmt.Printf("Name: %s, URL: %s",uploadedFile.Name , uploadedFile.URL)
	 }

}
