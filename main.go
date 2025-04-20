package main

import (
	"discordbot/configs"
	"discordbot/internals/utils"
	"fmt"
	"log"
	"time"
	"sync"
)

const (
	token = "<secret-token>"
)

func main() {
	var wg sync.WaitGroup
	configs := getConfigs()

	for _, config := range configs {
		wg.Add(1)
		go func ()  {
			defer wg.Done()	
			run(config)
		}()
	}

	wg.Wait()
}

func getConfigs() []configs.Config {
	res := []configs.Config{ }

	res = append(res, configs.EvoqConfigs...)

	res = append(res, configs.KuviConfigs...)

	return res
}

func run(config configs.Config) {
	retryCount := 0

	if retryCount > config.Retry {
		return
	}

	for {
		delay := utils.GetRandomDailyDelay()
		targetTime := time.Now().Add(delay)	

		log.Printf("🕒 Next message scheduled at: %s (in %s)\n", targetTime.Format(time.RFC1123), delay)

		time.Sleep(delay)


		message := utils.GetRandomMessage(config.Messages)
		resp, err := utils.SendDiscordMessage(utils.Message{
			Message: message,
			ChannelID: config.ChannelID,
			Token: token,
		})
	
		if err != nil {
			log.Fatalf("❌ Error sending message: %v", err)
			retryCount += 1
		}
	
		if resp.IsSuccess() {
			fmt.Println("✅ Message sent successfully!")
			retryCount = 0
		} else {
			fmt.Printf("❌ Failed to send message. Status: %v\n", resp.Status())
			fmt.Println("Response:", resp.String())
			retryCount += 1
		}
	}
}