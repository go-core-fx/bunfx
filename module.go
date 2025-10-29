package bunfx

import (
	"context"
	"database/sql"

	"github.com/go-core-fx/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"bun",
		logger.WithNamedLogger("bun"),
		fx.Provide(func(db *sql.DB, dialect schema.Dialect) *bun.DB {
			return bun.NewDB(db, dialect)
		}),
		fx.Invoke(func(db *bun.DB, lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					return nil
				},
				OnStop: func(_ context.Context) error {
					return db.Close()
				},
			})
		}),
	)
}
