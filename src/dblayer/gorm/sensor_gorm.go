package gorm

import "pdk/src/models"

func (db *GORM) GetAllSensors() (ss []models.Sensor, err error) {
	return ss, db.Find(&ss).Error
}

func (db *GORM) GetSensorsByNID(nID string) (sensors []models.Sensor, err error) {
	return sensors, db.Table("node_sensors").Select("*").Joins("join sensors on sensors.uuid=sensor_uuid").Where("node_uuid=?", nID).Scan(&sensors).Error
}

func (db *GORM) GetSensorValues(sID string) ([]string, error) {
	var svs []models.SensorValue
	if err := db.Table("sensor_values").Where("sensor_uuid=?", sID).Order("index").Scan(&svs).Error; err != nil {
		return nil, err
	}
	var vns = []string{}
	for _, sv := range svs {
		vns = append(vns, sv.ValueName)
	}
	return vns, nil
}

func (db *GORM) AddSensor(sensor models.Sensor) (models.Sensor, error) {
	return sensor, db.Create(&sensor).Error
}

func (db *GORM) AddSensorValue(sv models.SensorValue) (models.SensorValue, error) {
	return sv, db.Create(&sv).Error
}
