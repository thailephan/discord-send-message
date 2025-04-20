package utils

func GetRandomMessage(messages []string) string {
	messageSize := len(messages)
	message := messages[GenerateRandomIndex(messageSize)]
	return message
}