package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
)

type nodeRepo struct {
	db *gorm.DB
}

func NewNodeRepo() *nodeRepo {
	return &nodeRepo{
		db: dbConn,
	}
}

func (ndr *nodeRepo) FindsWithSensorsValues() (nl []model.Node, err error) {
	return nl, ndr.db.Preload("Sensors.SensorValues", orderByASC).Preload("Sensors").Find(&nl).Error
}

func (ndr *nodeRepo) Create(n *model.Node) error {
	return ndr.db.Omit("Sensors").Create(n).Association("Sensors").Append(n.Sensors)
}

func (ndr *nodeRepo) Delete(n *model.Node) error {
	return ndr.db.Delete(n).Error
}
