package registUsecase

import "github.com/KumKeeHyun/toiot/application/domain/model"

func (ru *registUsecase) GetLogics() ([]model.Logic, error) {
	return ru.lgr.FindsWithSensorValues()
}

func (ru *registUsecase) RegistLogic(l *model.Logic) error {
	return ru.lgr.Create(l)
}

func (ru *registUsecase) UnregistLogic(l *model.Logic) error {
	return ru.lgr.Delete(l)
}
