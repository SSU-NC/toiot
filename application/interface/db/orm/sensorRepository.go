package orm

import (
	"github.com/seheee/PDK/application/domain/model"
	"github.com/jinzhu/gorm"
)

type sensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository() *sensorRepository {
	return &sensorRepository{
		db: dbConn,
	}
}

func (sr *sensorRepository) GetAll() (s []model.Sensor, err error) {
	return s, sr.db.Find(&s).Error
}

func (sr *sensorRepository) GetByNodeUUID(nid string) (s []model.Sensor, err error) {
	return s, sr.db.Table("node_sensors").Select("*").Joins("join sensors on sensors.uuid=sensor_uuid").Where("node_uuid=?", nid).Scan(&s).Error

}

func (sr *sensorRepository) GetByUUID(sid string) (s *model.Sensor, err error) {
	return s, sr.db.Where("uuid=?", sid).Find(s).Error
}

func (sr *sensorRepository) GetValuesByUUID(sid string) (sv []model.SensorValue, err error) {
	return sv, sr.db.Table("sensor_values").Where("sensor_uuid=?", sid).Order("index").Scan(&sv).Error
}

func (sr *sensorRepository) Create(s *model.Sensor) error {
	return sr.db.Create(s).Error
}

func (sr *sensorRepository) Delete(s *model.Sensor) error {
	return sr.db.Delete(s).Error
}

func (sr *sensorRepository) CreateValue(sv *model.SensorValue) error {
	return sr.db.Create(sv).Error
}
