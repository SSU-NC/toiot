package setting

import (
	"log"
	"os"
)

type App struct {
	Server string
	React  string
}

func (as *App) Getenv() {
	as.Server = os.Getenv("APP_SERVER")
	if as.Server == "" {
		as.Server = "0.0.0.0:8081"
	}

	as.React = os.Getenv("REACT_SERVER")
	if as.React == "" {
		as.React = "220.70.2.171:3000"
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
		ds.Pass = "pdk1234"
	}
	ds.Database = os.Getenv("DB_DATABASE")
	if ds.Database == "" {
		ds.Database = "pdk"
	}
}

var Databasesetting = &Database{}

func init() {
	Appsetting.Getenv()
	Databasesetting.Getenv()

	log.Printf("app : %v\ndb : %v\n", Appsetting, Databasesetting)
}
