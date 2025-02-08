package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mimparat132/gonotify/pkg/discordwebhook"
)

type gonotifyFlags struct {
	username string
	content  string
}

func main() {

	flags := gonotifyFlags{}

	flag.StringVar(&flags.username, "username", "", "name of the service or user posting the message")
	flag.StringVar(&flags.content, "content", "", "message content to be sent to the discord server")

	flag.Parse()

	data, err := os.ReadFile("/etc/gonotify/gonotify_config.json")
	if err != nil {
		panic(err)
	}

	gonotifyConfig := discordwebhook.GonotifyConfig{}

	err = json.Unmarshal(data,&gonotifyConfig)
	if err != nil {
		fmt.Printf("could not get gonotify config: %v", err)
	}

	message := discordwebhook.Message{
		Username: &flags.username,
		Content:  &flags.content,
	}

	err = discordwebhook.SendMessage(gonotifyConfig.Webhookurl, message)
	if err != nil {
		fmt.Printf("could not send message to discord notification channel: %v\n", err)
	}

}
