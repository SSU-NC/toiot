package sql

import (
	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/jinzhu/gorm"
)

type sinkRepository struct {
	db *gorm.DB
}

func NewSinkRepository() *sinkRepository {
	return &sinkRepository{
		db: dbConn,
	}
}

func (sir *sinkRepository) GetAll() (sis []model.Sink, err error) {
	return sis, sir.db.Find(&sis).Error
}

func (sir *sinkRepository) GetAllWithNodes() (sis []model.Sink, err error) {
	return sis, sir.db.Table("sinks").Joins("join nodes on nodes.sink_id=id").Scan(&sis).Error
}

func (sir *sinkRepository) GetByID(id uint) (si *model.Sink, err error) {
	si = new(model.Sink)
	return si, sir.db.Where("id=?", id).First(si).Error
}

func (sir *sinkRepository) GetByIDWithNodes(id uint) (si *model.Sink, err error) {
	return si, sir.db.Table("sinks").Where("id=?", id).Joins("join nodes on nodes.sink_id=id").Scan(&si).Error
}

func (sir *sinkRepository) Create(si *model.Sink) error {
	return sir.db.Create(si).Error
}

func (sir *sinkRepository) Delete(si *model.Sink) error {
	return sir.db.Delete(si).Error
}
