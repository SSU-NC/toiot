package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type topicRepo struct {
	db *gorm.DB
}

func NewTopicRepo() *topicRepo {
	return &topicRepo{
		db: dbConn,
	}
}

func (tpr *topicRepo) FindsWithLogicService() (tl []model.Topic, err error) {
	return tl, tpr.db.Preload("LogicServices").Find(&tl).Error
}

func (tpr *topicRepo) Create(t *model.Topic) error {
	return tpr.db.Omit(clause.Associations).Create(t).Error
}

func (tpr *topicRepo) Delete(t *model.Topic) error {
	return tpr.db.Delete(t).Error
}
