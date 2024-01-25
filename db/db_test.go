package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gokch/cafe_manager/db/gen"
	"github.com/stretchr/testify/require"
)

var admin1 = gen.CreateAdminParams{
	ID:    "admin1",
	Name:  "admin1",
	Pw:    []byte("951357abcde!"),
	Phone: "010-1234-5678",
}

func TestAdmin(t *testing.T) {
	db, err := NewDB(ConnectFuncMysql("127.0.0.1", "3306", "root", "951753ck", "cafe"))
	if err != nil {
		t.Skip() // TODO : make db mockup server ( using docker testsuite )
	}
	defer db.Close()

	// create admin
	err = db.Job().CreateAdmin(context.Background(), admin1)
	require.NoError(t, err)

	// get exist admin
	adminInfo, err := db.Job().GetAdmin(context.Background(), "admin1")
	require.NoError(t, err)
	require.Equal(t, adminInfo.ID, admin1.ID)
	require.Equal(t, adminInfo.Name, admin1.Name)
	require.Equal(t, adminInfo.Phone, admin1.Phone)
	require.Equal(t, adminInfo.Pw, admin1.Pw)

	// get not exist admin
	adminInfo, err = db.Job().GetAdmin(context.Background(), "admin2")
	require.Equal(t, err, sql.ErrNoRows)
}

func TestTx(t *testing.T) {
}

func TestPrepareStatement(t *testing.T) {
}
