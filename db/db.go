package db

import (
	"database/sql"
)

func Connect(conn ConnectFunc) (db *Conn, err error) {
	db = &Conn{}
	db.DriverName, db.Dsn, db.DbName = conn()

	if db.db, err = sql.Open(db.DriverName, db.Dsn); err != nil {
		return nil, err
	}
	if err = db.db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

type Conn struct {
	DriverName string
	Dsn        string
	DbName     string

	db *sql.DB
}

func (t *Conn) Raw() *sql.DB {
	return t.db
}

func (t *Conn) SetOpenConns(openConns, idleConns int) {
	if openConns > 0 {
		t.db.SetMaxOpenConns(openConns)
	}
	if idleConns > 0 {
		t.db.SetMaxIdleConns(idleConns)
	}
}

func (t *Conn) Job() *Job {
	job := NewJob(t.db)
	return job
}

func (t *Conn) TxJob(isoLevel sql.IsolationLevel, readonly bool) (tx *Tx, err error) {
	tx, err = NewTx(t.db, isoLevel, readonly)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *Conn) TxJobFunc(isoLevel sql.IsolationLevel, readonly bool, fn func(*Tx) error) (err error) {
	tx, err := t.TxJob(isoLevel, readonly)
	if err != nil {
		return err
	}

	if err = fn(tx); err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
