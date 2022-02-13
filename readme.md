# *LAMBDA STARTER*

Lambda starter abstract the configurations of triggers in aws.

# *Triggers*

- [X] SQS
- [X] SNS
- [ ] S3
- [ ] Kinesis
- [ ] Api Gateway
- [ ] CloudWatch


# *How To Use*
To use is very simple you just need to implement the interface Handle and in your cmd call runner.Run passing your implementation of Handle.

```go
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
```
