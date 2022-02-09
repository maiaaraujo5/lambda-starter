package handler

import (
	"context"
	"github.com/maiaaraujo5/lambda-starter/event"
	"github.com/maiaaraujo5/lambda-starter/event/sns"
	"github.com/maiaaraujo5/lambda-starter/event/sqs"
)

type trigger func(event event.Event) []string

var triggers = map[string]trigger{
	"aws:sqs": sqs.Parse,
	"aws:sns": sns.Parse,
}

func Handle(ctx context.Context, event event.Event) error {

	messages := getMessages(event)

	return nil
}

func getMessages(event event.Event) []string {
	if len(event.Resources) == 0 {
		return nil
	}

	return triggers[event.Records[0].EventSource](event)
}
