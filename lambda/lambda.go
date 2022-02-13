package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/maiaaraujo5/lambda-starter/handler"
)

type Lambda struct {
	processer handler.Processer
}

func NewLambda(processer handler.Processer) *Lambda {
	return &Lambda{processer: processer}
}

func (h *Lambda) Start() {
	lambda.Start(h.processer.Process)
}
