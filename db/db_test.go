package db

import (
	"context"
	"testing"

	"github.com/gokch/cafe_manager/db/sqlc"
	"github.com/stretchr/testify/require"
)

// TODO
func TestExec(t *testing.T) {
	db, err := Connect(ConnectFuncMysql("127.0.0.1", "3306", "root", "951753ck", "cafe"))
	require.NoError(t, err)

	err = db.Job().Queries.CreateAdmin(context.Background(), sqlc.CreateAdminParams{
		ID:    "admin1",
		Name:  "admin1",
		Pw:    []byte("951357abcde!"),
		Phone: "010-1234-5678",
	})
	require.NoError(t, err)
}

func TestQuery(t *testing.T) {
}

func TestTx(t *testing.T) {
}

func TestPrepareStatement(t *testing.T) {
}
