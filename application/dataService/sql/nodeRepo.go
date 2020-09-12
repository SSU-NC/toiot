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
	return ndr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Sensors").Create(n).Error; err != nil {
			return err
		}
		s := n.Sensors
		n.Sensors = []model.Sensor{}
		if err := tx.Model(n).Association("Sensors").Append(s); err != nil {
			return err
		}
		return nil
	})
	// s := n.Sensors
	// n.Sensors = []model.Sensor{}
	// return ndr.db.Omit("Sensors").Create(n).Association("Sensors").Append(s)
}

func (ndr *nodeRepo) Delete(n *model.Node) error {
	return ndr.db.Delete(n).Error
}
