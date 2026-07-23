package go_pkg

import (
	"context"

	"github.com/google/uuid"
)

const UserIdKey = "UserId"

func PutUserId(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIdKey, userId)
}

func GetUserId(ctx context.Context) uuid.UUID {
	return ctx.Value(UserIdKey).(uuid.UUID)
}
