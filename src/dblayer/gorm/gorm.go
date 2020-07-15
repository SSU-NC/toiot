package gorm

import (
	"pdk/src/models"

	"github.com/jinzhu/gorm"
)

type GORM struct {
	*gorm.DB
}

var db GORM

func NewGORM(dbname, connect string) (*GORM, error) {
	db, err := gorm.Open(dbname, connect)
	db.AutoMigrate(
		&models.Node{},
		&models.Sensor{},
		&models.NodeSensor{},
		&models.SensorValue{},
	)
	db.Model(&models.NodeSensor{}).AddForeignKey("node_uuid", "nodes(uuid)", "CASCADE", "CASCADE")
	db.Model(&models.NodeSensor{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
	db.Model(&models.SensorValue{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")

	return &GORM{
		DB: db,
	}, err
}
