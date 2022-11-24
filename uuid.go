package sequence

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type UUID struct {
}

var (
	_ Sequence[string] = &UUID{}
)

func (g *UUID) Next(ctx context.Context) (string, error) {
	id := uuid.New()
	return id.String(), nil
}
