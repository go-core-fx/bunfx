package bunfx

import (
	"database/sql"

	"github.com/alexlast/bunzap"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"go.uber.org/zap"
)

func New(db *sql.DB, dialect schema.Dialect, logger *zap.Logger) *bun.DB {
	bundb := bun.NewDB(db, dialect)
	if logger.Level() == zap.DebugLevel {
		bundb.AddQueryHook(bunzap.NewQueryHook(bunzap.QueryHookOptions{
			Logger:       logger,
			SlowDuration: 0,
		}))
	}

	return bundb
}
