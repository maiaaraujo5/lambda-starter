package handler

import (
	"context"
	"github.com/maiaaraujo5/lambda-starter/event"
	"github.com/maiaaraujo5/lambda-starter/event/sns"
	"github.com/maiaaraujo5/lambda-starter/event/sqs"
)

type Processer interface {
	Process(ctx context.Context, event event.Event) error
}

type Process struct {
	handler Handler
}

type trigger func(event event.Event) []string

var triggers = map[string]trigger{
	"aws:sqs": sqs.Parse,
	"aws:sns": sns.Parse,
}

func NewProcess(handler Handler) *Process {
	return &Process{
		handler: handler,
	}
}

func (h *Process) Process(ctx context.Context, event event.Event) error {

	messages := h.getMessages(event)

	err := h.handler.Handle(ctx, messages)
	if err != nil {
		return err
	}

	return nil
}

func (h *Process) getMessages(event event.Event) []string {
	if len(event.Resources) == 0 {
		return nil
	}

	return triggers[event.Records[0].EventSource](event)
}
