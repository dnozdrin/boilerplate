package postgres

import (
	"context"
	"github.com/uptrace/bun"
)

type txKey struct{}

func injectTx(ctx context.Context, tx bun.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) (bun.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(bun.Tx)

	return tx, ok
}
