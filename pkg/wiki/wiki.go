package wiki

import (
	"context"
	"github.com/tomato3713/nullwiki/pkg/ent"
	"log/slog"
)

type Config struct {
	Logger   *slog.Logger
	Context  *context.Context
	DbClient *ent.Client
}
