package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type ConnectFunc func() (string, string, string)

// FIXME : is it neccessary check?
var (
	_ ConnectFunc = ConnectFuncMysql("", "", "", "", "")
	_ ConnectFunc = ConnectFuncPostgres("", "", "", "", "")
	_ ConnectFunc = ConnectFuncSqlite3("")
)

func ConnectFuncMysql(addr, port, id, pw, dbName string) ConnectFunc {
	return func() (string, string, string) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", id, pw, addr, port, dbName)
		return "mysql", dsn, dbName
	}
}

func ConnectFuncPostgres(addr, port, id, pw, dbName string) ConnectFunc {
	return func() (string, string, string) {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", addr, port, id, pw, dbName)
		return "postgres", dsn, dbName
	}
}

func ConnectFuncSqlite3(dbName string) ConnectFunc {
	return func() (string, string, string) {
		dsn := dbName
		return "sqlite3", dsn, dbName
	}
}
