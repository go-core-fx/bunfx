package bunfx

import "time"

type TimedModel struct {
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`
}
