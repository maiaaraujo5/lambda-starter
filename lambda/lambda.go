package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/maiaaraujo5/lambda-starter/handler"
)

type Lambda struct {
}

func NewLambda() *Lambda {
	return &Lambda{}
}

func (h *Lambda) Start() {
	lambda.Start(handler.Handle)
}
