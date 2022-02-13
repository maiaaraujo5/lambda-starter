package handler

import "context"

type Handler interface {
	Handle(ctx context.Context, message []string) error
}
