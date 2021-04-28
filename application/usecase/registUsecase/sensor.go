package registUsecase

import (
	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"log"
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
	log.Println("RegistSensor_s1 =",s)
	for i := range s.SensorValues {
		s.SensorValues[i].Index = i
	}
	log.Println("RegistSensor_s2 =",s)
	return ru.snr.Create(s)
}

func (ru *registUsecase) UnregistSensor(s *model.Sensor) error {
	return ru.snr.Delete(s)
}
