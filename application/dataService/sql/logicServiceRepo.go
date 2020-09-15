package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logicServiceRepo struct {
	db *gorm.DB
}

func NewLogicServiceRepo() *logicServiceRepo {
	return &logicServiceRepo{
		db: dbConn,
	}
}

func (lsr *logicServiceRepo) Finds() (ll []model.LogicService, err error) {
	return ll, lsr.db.Find(&ll).Error
}

func (lsr *logicServiceRepo) FindsWithTopic() (ll []model.LogicService, err error) {
	return ll, lsr.db.Preload("Topic").Find(&ll).Error
}

func (lsr *logicServiceRepo) FindsByTopicID(TopicID int) (ll []model.LogicService, err error) {
	return ll, lsr.db.Where("topic_id=?", TopicID).Find(&ll).Error
}

func (lsr *logicServiceRepo) FindByAddr(addr string) (l *model.LogicService, err error) {
	l = &model.LogicService{}
	return l, lsr.db.Where("addr=?", addr).Find(l).Error
}

func (lsr *logicServiceRepo) Create(l *model.LogicService) error {
	return lsr.db.Transaction(func(tx *gorm.DB) error {
		t := model.Topic{}
		if err := tx.Where("name=?", l.Topic.Name).Find(&t).Error; err != nil {
			return err
		}
		l.TopicID = t.ID
		if err := tx.Omit(clause.Associations).Create(l).Error; err != nil {
			return err
		}
		return nil
	})
	// return lsr.db.Omit(clause.Associations).Create(l).Error
}

func (lsr *logicServiceRepo) Delete(l *model.LogicService) error {
	return lsr.db.Delete(l).Error
}
