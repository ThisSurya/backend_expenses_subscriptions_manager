package db

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func NewDb(db *gorm.DB) *handler {
	return &handler{DB: db}
}
