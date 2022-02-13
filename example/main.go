package main

import (
	"context"
	"github.com/maiaaraujo5/lambda-starter/handler"
	lambdaRunner "github.com/maiaaraujo5/lambda-starter/runner"
	"log"
)

func main() {
	e := newExample()

	err := lambdaRunner.Run(e)
	if err != nil {
		panic(err)
	}
}

type example struct {
}

func newExample() handler.Handler {
	return &example{}
}

func (h *example) Handle(ctx context.Context, messages []string) error {
	log.Println("running lambda")
	for _, message := range messages {
		log.Println(message)
	}

	return nil
}
