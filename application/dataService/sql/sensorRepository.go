package sql

import (
	"github.com/KumKeeHyun/PDK/application/domain/model"
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

func (sr *sensorRepository) GetAll() (ss []model.Sensor, err error) {
	return ss, sr.db.Find(&ss).Error
}

func (sr *sensorRepository) GetByNodeUUID(nid string) (ss []model.Sensor, err error) {
	return ss, sr.db.Table("node_sensors").Select("*").Joins("join sensors on sensors.uuid=sensor_uuid").Where("node_uuid=?", nid).Scan(&ss).Error
}

func (sr *sensorRepository) GetByNodeUUIDWithValues(nid string) (ss []model.Sensor, err error) {
	err = sr.db.Table("node_sensors").Select("*").Joins("join sensors on sensors.uuid=sensor_uuid").Where("node_uuid=?", nid).Scan(&ss).Error
	if err != nil {
		return nil, err
	}

	for i := range ss {
		if ss[i].ValueList, err = sr.GetValuesByUUID(ss[i].UUID); err != nil {
			return nil, err
		}
	}
	return ss, nil
}

func (sr *sensorRepository) GetByUUID(sid string) (s *model.Sensor, err error) {
	return s, sr.db.Where("uuid=?", sid).Find(s).Error
}

func (sr *sensorRepository) GetByUUIDWithValues(sid string) (s *model.Sensor, err error) {
	err = sr.db.Table("sensors").Where("uuid=?", sid).Scan(s).Error
	if err != nil {
		return nil, err
	}
	if s.ValueList, err = sr.GetValuesByUUID(sid); err != nil {
		return nil, err
	}
	return s, nil

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
