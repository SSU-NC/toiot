package sensorUsecase

import (
	"github.com/seheee/PDK/application/domain/model"
	"github.com/seheee/PDK/application/domain/repository"
)

type sensorUsecase struct {
	sr repository.SensorRepository
}

func NewSensorUsecase(sr repository.SensorRepository) *sensorUsecase {
	return &sensorUsecase{
		sr: sr,
	}
}

func (su *sensorUsecase) GetAllSensors() ([]model.Sensor, error) {
	ss, err := su.sr.GetAll()
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (su *sensorUsecase) GetAllSensorsWithValues() ([]model.Sensor, error) {
	ss, err := su.sr.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range ss {
		if ss[i].ValueList, err = su.sr.GetValuesByUUID(ss[i].UUID); err != nil {
			return nil, err
		}
	}
	return ss, nil

}

func (su *sensorUsecase) RegisterSensor(s *model.Sensor) (*model.Sensor, error) {
	newSensor := model.NewSensor(s.Name)
	if err := su.sr.Create(&newSensor); err != nil {
		return nil, err
	}
	for i, v := range s.ValueList {
		v.SensorUUID = newSensor.UUID
		v.Index = i
		if err := su.sr.CreateValue(&v); err != nil {
			return nil, err
		}
		newSensor.ValueList = append(newSensor.ValueList, v)
	}
	return &newSensor, nil
}

func (su *sensorUsecase) DeleteSensor(s *model.Sensor) (*model.Sensor, error) {
	if err := su.sr.Delete(s); err != nil {
		return nil, err
	}
	return s, nil
}
