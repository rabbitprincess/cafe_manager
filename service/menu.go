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

func (m *Menu) GetMenu(seq int64) (*gen.Product, error) {
	product, err := m.db.Job().GetProduct(context.Background(), seq)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (m *Menu) ListMenu(seq int64) ([]*gen.Product, error) {
	products, err := m.db.Job().ListProducts(context.Background(), seq)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return products, nil
}

func (m *Menu) SearchMenu(name string) ([]*gen.Product, error) {
	var products []*gen.Product
	var err error
	if utilx.IsHangulInitialsOnly(name) {
		products, err = m.db.Job().SearchProductsByNameInitial(context.Background(), gen.SearchProductsByNameInitialParams{
			NameInitial: name,
			Limit:       10,
		})
	} else {
		products, err = m.db.Job().SearchProductsByName(context.Background(), gen.SearchProductsByNameParams{
			Name:  name,
			Limit: 10,
		})
	}
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return products, nil
}

func (m *Menu) AddMenu(category, name, description string, price, cost, expire int64, size bool) error {
	barcode, err := utilx.Barcode(name)
	if err != nil {
		return fmt.Errorf("failed to generate barcode from name | name : %s | err : %w", name, err)
	}
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.CreateProduct(context.Background(), gen.CreateProductParams{
			Category:    sql.NullString{String: category, Valid: true},
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

func (m *Menu) UpdateMenu(category, name, description *string, price, cost, expire *int64, size *bool) error {
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
	var sizeNull sql.NullBool
	if size != nil {
		sizeNull = sql.NullBool{Bool: *size, Valid: true}
	}
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.UpdateProductIfNotNil(context.Background(), gen.UpdateProductIfNotNilParams{
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

func (m *Menu) DeleteMenu(seq int64) error {
	return m.db.TxJobFunc(sql.LevelDefault, false, func(tx *db.Tx) error {
		if err := tx.DeleteProduct(context.Background(), seq); err != nil {
			return err
		}
		return nil
	})
}
