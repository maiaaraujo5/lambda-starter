package event

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type Record struct {
	EventVersion         string           `json:"eventVersion"`
	EventSubscriptionArn string           `json:"eventSubscriptionArn"`
	EventSource          string           `json:"eventSource"`
	EventName            string           `json:"eventName"`
	EventID              string           `json:"eventID"`
	SNS                  events.SNSEntity `json:"sns"`
	events.SQSMessage
}

type Event struct {
	ID         string    `json:"id"`
	Source     string    `json:"source"`
	Region     string    `json:"region"`
	DetailType string    `json:"detail-type"`
	Time       time.Time `json:"time"`
	Account    string    `json:"account"`
	Resources  []string  `json:"resources"`
	Records    []Record  `json:"Records"`
}
