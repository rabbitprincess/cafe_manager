package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gokch/cafe_manager/db/gen"
)

// non-tx job
func NewJob(db *sql.DB) *Job {
	return &Job{
		Queries: gen.New(db),
		db:      db,
	}
}

type Job struct {
	*gen.Queries
	db *sql.DB
}

// tx job
func NewTx(db *sql.DB, isoLevel sql.IsolationLevel, readOnly bool) (*Tx, error) {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: isoLevel,
		ReadOnly:  readOnly,
	})
	if err != nil {
		return nil, err
	}
	return &Tx{
		Queries: gen.New(tx),
		tx:      tx,
	}, nil
}

type Tx struct {
	closed bool
	*gen.Queries
	tx *sql.Tx
}

func (t *Tx) Commit() error {
	if t.closed {
		return errors.New("tx already commit or rollback")
	}
	if err := t.tx.Commit(); err != nil {
		t.tx.Rollback()
		return err
	}
	t.closed = true
	return nil
}

func (t *Tx) Rollback() error {
	if t.closed {
		return errors.New("tx already commit or rollback")
	}
	t.tx.Rollback()
	t.closed = true
	return nil
}
