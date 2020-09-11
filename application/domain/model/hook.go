package model

import (
	"fmt"

	"gorm.io/gorm"
)

var orderByASC = func(db *gorm.DB) *gorm.DB {
	return db.Order("sensor_values.index ASC")
}

func (s *Sink) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("Topic").Find(s).Error
}

func (s *Sink) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Topic").Preload("Nodes").Find(s).Error
}

func (n *Node) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("Sink.Topic").Preload("Sink").Find(n).Error
}

func (n *Node) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Sink.Topic").Preload("Sink").Find(n).Error
}

func (s *Sensor) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("SensorValues", orderByASC).Find(s).Error
}

func (s *Sensor) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Nodes.Sink.Topic").Preload("Nodes.Sink").Preload("Nodes").Find(s).Error
}

func (t *Topic) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Preload("LogicServices").Find(t).Error; err != nil {
		return err
	}
	if len(t.LogicServices) != 0 {
		return fmt.Errorf("There are logic-services that consume topic : %s", t.Name)
	}
	return nil
}
