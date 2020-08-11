package sql

import (
	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/jinzhu/gorm"
)

type nodeRepository struct {
	db *gorm.DB
}

func NewNodeRepository() *nodeRepository {
	return &nodeRepository{
		db: dbConn,
	}
}

func (nr *nodeRepository) GetAll() (ns []model.Node, err error) {
	return ns, nr.db.Find(&ns).Error
}

func (nr *nodeRepository) GetByUUID(nid string) (n *model.Node, err error) {
	return n, nr.db.Where("uuid=?", nid).Find(n).Error
}

func (nr *nodeRepository) GetBySinkID(sinkID uint) (ns []model.Node, err error) {
	return ns, nr.db.Where("sink_id=?", sinkID).Find(&ns).Error
}

func (nr *nodeRepository) Create(n *model.Node) error {
	return nr.db.Create(n).Error
}

func (nr *nodeRepository) Delete(n *model.Node) error {
	return nr.db.Delete(n).Error
}

func (nr *nodeRepository) CreateNS(ns *model.NodeSensor) error {
	return nr.db.Create(ns).Error
}
