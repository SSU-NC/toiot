package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
)

var orderByASC = func(db *gorm.DB) *gorm.DB {
	return db.Order("sensor_values.index ASC")
}

type sensorRepo struct {
	db *gorm.DB
}

func NewSensorRepo() *sensorRepo {
	return &sensorRepo{
		db: dbConn,
	}
}

func (snr *sensorRepo) FindsWithValues() (sl []model.Sensor, err error) {
	return sl, snr.db.Preload("SensorValues", orderByASC).Find(&sl).Error
}

func (snr *sensorRepo) Create(s *model.Sensor) error {
	return snr.db.Create(s).Error
}

func (snr *sensorRepo) Delete(s *model.Sensor) error {
	return snr.db.Delete(s).Error
}
