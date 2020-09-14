package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logicRepo struct {
	db *gorm.DB
}

func NewLogicRepo() *logicRepo {
	return &logicRepo{
		db: dbConn,
	}
}

func (lgr *logicRepo) FindsWithSensorValues() (ll []model.Logic, err error) {
	return ll, lgr.db.Preload("Sensor").Find(&ll).Error
}

func (lgr *logicRepo) Create(l *model.Logic) error {
	return lgr.db.Omit(clause.Associations).Create(l).Error
}

func (lgr *logicRepo) Delete(l *model.Logic) error {
	return lgr.db.Delete(l).Error
}
