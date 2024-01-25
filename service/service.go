package service

import "github.com/gokch/cafe_manager/db"

func NewService(db *db.DB) *Service {
	return &Service{
		db:    db,
		Admin: &Admin{db: db},
		Menu:  &Menu{db: db},
	}
}

type Service struct {
	db    *db.DB
	Admin *Admin
	Menu  *Menu
}
