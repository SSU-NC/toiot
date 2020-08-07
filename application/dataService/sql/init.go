package sql

import (
	"fmt"

	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/KumKeeHyun/PDK/application/setting"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

func Setup() {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", setting.Databasesetting.User, setting.Databasesetting.Pass, setting.Databasesetting.TCP, setting.Databasesetting.Database)
	dbConn, _ = gorm.Open(setting.Databasesetting.Driver, conn)

	dbConn.AutoMigrate(
		&model.Node{},
		&model.Sensor{},
		&model.NodeSensor{},
		&model.SensorValue{},
	)
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("node_uuid", "nodes(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.SensorValue{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
}
