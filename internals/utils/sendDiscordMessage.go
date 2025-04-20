package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Message struct {
	Token string
	Message string
	ChannelID string
}
func SendDiscordMessage(args Message) (*resty.Response, error){
	url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", args.ChannelID)

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", args.Token).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"content": args.Message,
			"mobile_network_type": "unknown",
			"tts": "false",
			"flags": "0",
		}).
		Post(url)
	return resp, err
}