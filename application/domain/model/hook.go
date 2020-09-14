package model

import (
	"fmt"

	"gorm.io/gorm"
)

var orderByASC = func(db *gorm.DB) *gorm.DB {
	return db.Order("sensor_values.index ASC")
}

// sink
func (s *Sink) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("Topic").Find(s).Error
}

func (s *Sink) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Topic").Preload("Nodes").Find(s).Error
}

// node
func (n *Node) AfterCreate(tx *gorm.DB) (err error) {
	if err := tx.Preload("Sink.Topic").Preload("Sink").Find(n).Error; err != nil {
		return err
	}
	l := []Logic{}
	// TODO : I want to Preload("Sensors.Logics") but it dosen't work for me :(
	for i := range n.Sensors {
		if err := tx.Where("sensor_id=?", n.Sensors[i].ID).Find(&l).Error; err != nil {
			return err
		}
		n.Sensors[i].Logics = l
	}
	return nil
}

func (n *Node) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Sink.Topic").Preload("Sink").Find(n).Error
}

// sensor
func (s *Sensor) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("SensorValues", orderByASC).Find(s).Error
}

func (s *Sensor) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Logics").Find(s).Error
}

// logic
func (l *Logic) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("Sensor").Find(l).Error
}

func (l *Logic) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Sensor").Find(l).Error
}

// logicService
func (l *LogicService) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Preload("Topic.Sinks.Nodes.Sensors.Logics").Preload("Topic.Sinks.Nodes.Sensors.SensorValues", orderByASC).Preload("Topic.Sinks.Nodes.Sensors").Preload("Topic.Sinks.Nodes").Preload("Topic.Sinks").Preload("Topic").Find(l).Error
}

func (l *LogicService) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Preload("Topic").Find(l).Error
}

// topic
func (t *Topic) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Preload("LogicServices").Find(t).Error; err != nil {
		return err
	}
	if len(t.LogicServices) != 0 {
		return fmt.Errorf("There are logic-services that consume topic : %s", t.Name)
	}
	return nil
}
