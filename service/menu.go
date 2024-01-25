package service

import (
	"context"

	"github.com/gokch/cafe_manager/db"
	"github.com/gokch/cafe_manager/db/gen"
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
