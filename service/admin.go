package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gokch/cafe_manager/db"
	"github.com/gokch/cafe_manager/db/sqlc"
	"github.com/gokch/cafe_manager/utilx"
)

type Admin struct {
	db *db.DB
}

func (a *Admin) Register(id, name, pw, phone string) error {
	pwSecured, err := utilx.BEncrypt(pw)
	if err != nil {
		return err
	}
	if !utilx.IsPhoneNumber(phone) {
		return fmt.Errorf("invalid phone number: %s", phone)
	}

	return a.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err = tx.CreateAdmin(context.Background(), sqlc.CreateAdminParams{
			ID:    id,
			Name:  name,
			Pw:    pwSecured,
			Phone: phone,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (a *Admin) Login() {

}

func (a *Admin) Logout() {

}
