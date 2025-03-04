package port

import "context"

type IMeatRepository interface {
	Get(ctx context.Context, index string) (string, error)
}
