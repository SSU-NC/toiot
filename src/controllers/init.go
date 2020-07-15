package controllers

import (
	"pdk/src/dblayer"
	"pdk/src/dblayer/gorm"
	"pdk/src/setting"
)

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	d, cn := setting.Databasesetting.MakeConnection()
	db, err := gorm.NewGORM(d, cn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
