package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gokch/cafe_manager/db"
	"github.com/gokch/cafe_manager/db/gen"
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
		if err = tx.CreateAdmin(context.Background(), gen.CreateAdminParams{
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

func (a *Admin) Login(id, pw string) error {
	adminInfo, err := a.db.Job().GetAdmin(context.Background(), id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("invalid id")
	} else if err != nil {
		return err
	}
	pwValid, err := utilx.BCheck(pw, adminInfo.Pw)
	if err != nil {
		return err
	} else if pwValid != true {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (a *Admin) Logout() {

}
