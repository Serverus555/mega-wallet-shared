package shared

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

const TraceIdKey = "trace_id"
const UserIdKey = "user_id"

func PutUserId(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIdKey, userId)
}

func GetUserId(ctx context.Context) uuid.UUID {
	// panic специально
	return ctx.Value(UserIdKey).(uuid.UUID)
}

func RichLogger(ctx context.Context, logger zerolog.Logger) zerolog.Logger {
	return logger.With().
		Str(TraceIdKey, fmt.Sprintf("%v", ctx.Value(TraceIdKey))).
		Str(UserIdKey, fmt.Sprintf("%v", ctx.Value(UserIdKey))).
		Logger()
}
