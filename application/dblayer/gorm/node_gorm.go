package gorm

import "pdk/src/models"

func (db *GORM) GetAllNodes() (nodes []models.Node, err error) {
	return nodes, db.Find(&nodes).Error
}

func (db *GORM) AddNode(node models.Node) (models.Node, error) {
	return node, db.Create(&node).Error
}

func (db *GORM) AddNodeSensor(ns models.NodeSensor) (models.NodeSensor, error) {
	return ns, db.Create(&ns).Error
}
