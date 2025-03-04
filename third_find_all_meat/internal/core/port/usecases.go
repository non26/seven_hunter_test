package port

import "context"

type IMeatUsecase interface {
	GetAll(ctx context.Context) (map[string]int, error)
}
