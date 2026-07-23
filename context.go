package go_pkg

import (
	"context"

	"github.com/google/uuid"
)

type userIdKey struct{}

func PutUserId(ctx context.Context, userId uuid.UUID) {
	context.WithValue(ctx, userIdKey{}, userId)
}

func GetUserId(ctx context.Context) uuid.UUID {
	return ctx.Value(userIdKey{}).(uuid.UUID)
}
