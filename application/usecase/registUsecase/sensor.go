package registUsecase

import (
	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

func (ru *registUsecase) GetSensorPageCount(size int) int {
	return ru.snr.GetPages(size)
}

func (ru *registUsecase) GetSensors() ([]model.Sensor, error) {
	return ru.snr.FindsWithValues()
}

func (ru *registUsecase) GetSensorsPage(p adapter.Page) ([]model.Sensor, error) {
	return ru.snr.FindsPage(p)
}

func (ru *registUsecase) RegistSensor(s *model.Sensor) error {
	for i := range s.SensorValues {
		s.SensorValues[i].Index = i
	}
	return ru.snr.Create(s)
}

func (ru *registUsecase) UnregistSensor(s *model.Sensor) error {
	return ru.snr.Delete(s)
}
