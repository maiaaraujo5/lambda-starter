package sns

import "github.com/maiaaraujo5/lambda-starter/event"

func Parse(event event.Event) []string {
	var messages []string
	for _, record := range event.Records {
		messages = append(messages, record.SNS.Message)
	}
	return messages
}
