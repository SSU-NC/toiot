package setting

import (
	"log"
	"os"
	"strconv"
)

func GetenvInt(target *int, init int, env string) {
	var err error

	temp := os.Getenv(env)
	if temp == "" {
		*target = init
	} else {
		if *target, err = strconv.Atoi(temp); err != nil {
			*target = init
		}
	}
}

type App struct {
	Server string
}

func (as *App) Getenv() {
	as.Server = os.Getenv("APP_SERVER")
	if as.Server == "" {
		as.Server = "10.5.110.11:8081"
	}
}

var Appsetting = &App{}

type Database struct {
	Driver   string `toml:"driver"`
	Server   string `toml:"tcp"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	Database string `toml:"database"`
}

func (ds *Database) Getenv() {
	ds.Driver = os.Getenv("DB_DRIVER")
	if ds.Driver == "" {
		ds.Driver = "mysql"
	}
	ds.Server = os.Getenv("DB_SERVER")
	if ds.Server == "" {
		ds.Server = "localhost:3306"
	}
	ds.User = os.Getenv("DB_USER")
	if ds.User == "" {
		ds.User = "pdk"
	}
	ds.Pass = os.Getenv("DB_PASS")
	if ds.Pass == "" {
		ds.Pass = "pdk12345"
	}
	ds.Database = os.Getenv("DB_DATABASE")
	if ds.Database == "" {
		ds.Database = "pdk"
	}
}

var Databasesetting = &Database{}

type Topic struct {
	Name         string
	Partitions   int
	Replications int
}

func (ts *Topic) Getenv() {
	ts.Name = os.Getenv("TOPIC_NAME")
	if ts.Name == "" {
		ts.Name = "sensor-data"
	}
	GetenvInt(&ts.Partitions, 1, "TOPIC_PARTITIONS")
	GetenvInt(&ts.Replications, 1, "TOPIC_REPLICATIONS")
}

var Topicsetting = &Topic{}

func init() {
	Appsetting.Getenv()
	Databasesetting.Getenv()
	Topicsetting.Getenv()

	log.Printf("app : %v\ndb : %v\n", Appsetting, Databasesetting)
}
