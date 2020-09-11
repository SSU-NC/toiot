package registUsecase

import "github.com/KumKeeHyun/toiot/application/domain/model"

func (ru *registUsecase) GetSensors() ([]model.Sensor, error) {
	return ru.snr.FindsWithValues()
}

func (ru *registUsecase) RegistSensor(s *model.Sensor) error {
	return ru.snr.Create(s)
}

func (ru *registUsecase) UnregistSensor(s *model.Sensor) error {
	return ru.snr.Delete(s)
}
