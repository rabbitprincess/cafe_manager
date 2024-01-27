package service

import (
	"context"
	"database/sql"
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

func (m *Menu) AddMenu(category, name, description string, price, cost, expire int64, barcode, size string) error {
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.CreateMenu(context.Background(), gen.CreateMenuParams{
			Category:    category,
			Price:       int32(price),
			Cost:        int32(cost),
			Name:        name,
			Description: description,
			Expire:      time.Unix(expire, 0),
			Barcode:     barcode,
			Size:        size,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (m *Menu) UpdateMenu(category, name, description string, price, cost, expire int64, barcode, size string) error {
	var categoryNull, nameNull, decriptionNull sql.NullString
	if category != "" {
		categoryNull = sql.NullString{String: category, Valid: true}
	}
	if name != "" {
		nameNull = sql.NullString{String: name, Valid: true}
	}
	if description != "" {
		decriptionNull = sql.NullString{String: description, Valid: true}
	}

	var priceNull, costNull sql.NullInt32
	if price != 0 {
		priceNull = sql.NullInt32{Int32: int32(price), Valid: true}
	}
	if cost != 0 {
		costNull = sql.NullInt32{Int32: int32(cost), Valid: true}
	}

	var expireNull sql.NullTime
	if expire != 0 {
		expireNull = sql.NullTime{Time: time.Unix(expire, 0), Valid: true}
	}
	var barcodeNull, sizeNull sql.NullString
	if barcode != "" {
		barcodeNull = sql.NullString{String: barcode, Valid: true}
	}
	if size != "" {
		sizeNull = sql.NullString{String: size, Valid: true}
	}
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.UpdateMenuIfNotNil(context.Background(), gen.UpdateMenuIfNotNilParams{
			Category:   categoryNull,
			Price:      priceNull,
			Cost:       costNull,
			Name:       nameNull,
			Decription: decriptionNull,
			Expire:     expireNull,
			Barcode:    barcodeNull,
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
