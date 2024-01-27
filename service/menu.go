package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gokch/cafe_manager/db"
	"github.com/gokch/cafe_manager/db/gen"
	"github.com/gokch/cafe_manager/utilx"
)

type Menu struct {
	db *db.DB
}

func (m *Menu) GetMenu(seq uint64) (*gen.Menu, error) {
	menu, err := m.db.Job().GetMenu(context.Background(), seq)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (m *Menu) ListMenu(seq uint64) ([]*gen.Menu, error) {
	menus, err := m.db.Job().ListMenus(context.Background(), seq)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *Menu) SearchMenu(name string) ([]*gen.Menu, error) {
	var menus []*gen.Menu
	var err error
	if utilx.IsHangulInitialsOnly(name) {
		menus, err = m.db.Job().SearchMenusByNameInitial(context.Background(), gen.SearchMenusByNameInitialParams{
			NameInitial: name,
			Limit:       10,
		})
	} else {
		menus, err = m.db.Job().SearchMenusByName(context.Background(), gen.SearchMenusByNameParams{
			Name:  name,
			Limit: 10,
		})
	}
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *Menu) AddMenu(category, name, description string, price, cost, expire int64, size string) error {
	barcode, err := utilx.Barcode(name)
	if err != nil {
		return fmt.Errorf("failed to generate barcode from name | name : %s | err : %w", name, err)
	}
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.CreateMenu(context.Background(), gen.CreateMenuParams{
			Category:    category,
			Price:       int32(price),
			Cost:        int32(cost),
			Name:        name,
			Description: description,
			Barcode:     []byte(barcode.Content()),
			Expire:      time.Unix(expire, 0),
			Size:        size,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (m *Menu) UpdateMenu(category, name, description *string, price, cost, expire *int64, size *string) error {
	var categoryNull, nameNull, decriptionNull sql.NullString
	if category != nil {
		categoryNull = sql.NullString{String: *category, Valid: true}
	}
	if name != nil {
		nameNull = sql.NullString{String: *name, Valid: true}
	}
	if description != nil {
		decriptionNull = sql.NullString{String: *description, Valid: true}
	}

	var priceNull, costNull sql.NullInt32
	if price != nil {
		priceNull = sql.NullInt32{Int32: int32(*price), Valid: true}
	}
	if cost != nil {
		costNull = sql.NullInt32{Int32: int32(*cost), Valid: true}
	}

	var expireNull sql.NullTime
	if expire != nil {
		expireNull = sql.NullTime{Time: time.Unix(*expire, 0), Valid: true}
	}
	var sizeNull sql.NullString
	if size != nil {
		sizeNull = sql.NullString{String: *size, Valid: true}
	}
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.UpdateMenuIfNotNil(context.Background(), gen.UpdateMenuIfNotNilParams{
			Category:   categoryNull,
			Price:      priceNull,
			Cost:       costNull,
			Name:       nameNull,
			Decription: decriptionNull,
			Expire:     expireNull,
			Size:       sizeNull,
		}); err != nil {
			return err
		}
		return nil
	})

}

func (m *Menu) DeleteMenu(seq uint64) error {
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.DeleteMenu(context.Background(), seq); err != nil {
			return err
		}
		return nil
	})
}
