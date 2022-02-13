package runner

import (
	"github.com/maiaaraujo5/lambda-starter/handler"
	"github.com/maiaaraujo5/lambda-starter/lambda"
)

func Run(h handler.Handler) error {
	lambda.NewLambda(
		handler.NewProcess(h),
	).Start()

	return nil
}
