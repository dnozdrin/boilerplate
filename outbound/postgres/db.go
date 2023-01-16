package postgres

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type DB struct {
	conn bun.IDB
	log  ErrorLogger
}

// ErrorLogger - WIP
// TODO: review interface and usage
type ErrorLogger interface {
	Error(msg string, args ...interface{})
}

func NewDB(db bun.IDB) *DB {
	return &DB{
		conn: db,
	}
}

func (db *DB) WithConn(ctx context.Context) bun.IDB {
	tx, exist := extractTx(ctx)
	if exist {
		return tx
	}

	return db.conn
}

func (db *DB) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, beginErr := db.conn.BeginTx(ctx, nil)
	if beginErr != nil {
		return fmt.Errorf("begin transaction: %w", beginErr)
	}

	var done bool
	defer func() {
		if !done {
			if errRollback := tx.Rollback(); errRollback != nil {
				db.log.Error("rollback transaction: %v", errRollback)
			}
		}
	}()

	if funcErr := tFunc(injectTx(ctx, tx)); funcErr != nil {
		return funcErr
	}

	done = true
	if errCommit := tx.Commit(); errCommit != nil {
		db.log.Error("commit transaction: %v", errCommit)

		return errCommit
	}

	return nil
}
