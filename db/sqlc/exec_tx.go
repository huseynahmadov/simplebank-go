package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, txFunc func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		return err
	}

	q := New(tx)
	err = txFunc(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}