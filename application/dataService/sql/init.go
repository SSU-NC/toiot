package sql

import (
	"fmt"
	"log"
	"time"

	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/KumKeeHyun/PDK/application/setting"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

func Setup() {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", setting.Databasesetting.User, setting.Databasesetting.Pass, setting.Databasesetting.Server, setting.Databasesetting.Database)
	dbConn, _ = gorm.Open(setting.Databasesetting.Driver, conn)

	retry := 30
	for {
		err := dbConn.DB().Ping()
		if err != nil {
			dbConn, _ = gorm.Open(setting.Databasesetting.Driver, conn)
			if retry == 0 {
				log.Fatalf("Not able to establish connection to database")
			}
			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retry))
			retry--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	dbConn.AutoMigrate(
		&model.Sink{},
		&model.Node{},
		&model.Sensor{},
		&model.NodeSensor{},
		&model.SensorValue{},
	)
	dbConn.Model(&model.Node{}).AddForeignKey("sink_id", "sinks(id)", "CASCADE", "CASCADE")
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("node_uuid", "nodes(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.SensorValue{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
}
