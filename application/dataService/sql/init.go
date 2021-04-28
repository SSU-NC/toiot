package sql

import (
	"errors"
	"fmt"

	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/KumKeeHyun/toiot/application/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Setup() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", setting.Databasesetting.User, setting.Databasesetting.Pass, setting.Databasesetting.Server, setting.Databasesetting.Database)
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(errors.New("DB connection fail"))
	}

	dbConn.AutoMigrate(
		&model.Topic{}, &model.LogicService{},
		&model.Sink{}, &model.Node{},
		&model.Sensor{}, &model.SensorValue{}, &model.Logic{}, &model.Actuator{},
	)

	// dbConn.Model(&model.LogicService{}).AddForeignKey("topic_id", "topics(id)", "CASCADE", "CASCADE")
	// dbConn.Model(&model.Sink{}).AddForeignKey("topic_id", "topics(id)", "CASCADE", "CASCADE")
	// dbConn.Model(&model.Node{}).AddForeignKey("sink_id", "sinks(id)", "CASCADE", "CASCADE")
	// dbConn.Model(&model.SensorValue{}).AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE")
	// dbConn.Model(&model.Logic{}).AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE")
}
