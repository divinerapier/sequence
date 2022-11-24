package sequence

import "context"

type Sequence[T any] interface {
	Next(ctx context.Context) (T, error)
}
