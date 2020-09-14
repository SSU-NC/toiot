package sql

import (
	"github.com/KumKeeHyun/toiot/application/adapter"
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

func (ndr *nodeRepo) GetPages(p adapter.Page) int {
	temp := []model.Node{}
	if p.Sink != 0 {
		result := ndr.db.Where("sink_id=?", p.Sink).Find(&temp)
		count := int(result.RowsAffected)
		return (count / p.Size) + 1
	} else {
		result := ndr.db.Find(&temp)
		count := int(result.RowsAffected)
		return (count / p.Size) + 1
	}
}

func (ndr *nodeRepo) FindsWithSensorsValues() (nl []model.Node, err error) {
	return nl, ndr.db.Preload("Sensors.SensorValues", orderByASC).Preload("Sensors").Find(&nl).Error
}

func (ndr *nodeRepo) FindsPage(p adapter.Page) (nl []model.Node, err error) {
	offset := p.GetOffset()
	if p.Sink == 0 {
		return nl, ndr.db.Offset(offset).Limit(p.Size).Preload("Sensors.SensorValues", orderByASC).Preload("Sensors").Find(&nl).Error
	} else {
		return nl, ndr.db.Where("sink_id=?", p.Sink).Offset(offset).Limit(p.Size).Preload("Sensors.SensorValues", orderByASC).Preload("Sensors").Find(&nl).Error
	}
}

func (ndr *nodeRepo) FindsSquare(sq adapter.Square) (nl []model.Node, err error) {
	return nl, ndr.db.Where("loc_lon BETWEEN ? AND ?", sq.Left, sq.Right).Where("loc_lat BETWEEN ? AND ?", sq.Down, sq.Up).Preload("Sensors.SensorValues", orderByASC).Preload("Sensors").Find(&nl).Error
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
