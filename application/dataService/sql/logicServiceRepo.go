package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logicServiceRepo struct {
	db *gorm.DB
}

func NewLogicService() *logicServiceRepo {
	return &logicServiceRepo{
		db: dbConn,
	}
}

func (lsr *logicServiceRepo) FindsWithTopic() (ll []model.LogicService, err error) {
	return ll, lsr.db.Preload("Topic").Find(&ll).Error
}

func (lsr *logicServiceRepo) Create(l *model.LogicService) error {
	return lsr.db.Omit(clause.Associations).Create(l).Error
}

func (lsr *logicServiceRepo) Delete(l *model.LogicService) error {
	return lsr.db.Delete(l).Error
}
