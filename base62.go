package sequence

import (
	"context"

	"github.com/divinerapier/base62"
)

type Base62 struct {
	inner Sequence[uint64]
}

var (
	_ Sequence[string] = &Base62{}
)

func NewBase62(inner Sequence[uint64]) *Base62 {
	return &Base62{inner: inner}
}

func (s *Base62) Next(ctx context.Context) (string, error) {
	seq, err := s.inner.Next(ctx)
	if err != nil {
		return "", err
	}
	return base62.Encode(seq), nil
}
