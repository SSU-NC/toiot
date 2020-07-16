package setting

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

type Server struct {
	Address string
	Port    string
}

func (c *Server) MakeAddr() string {
	return fmt.Sprintf("%s:%s", c.Address, c.Port)
}

var Serversetting = &Server{}

type Database struct {
	Driver   string `toml:"driver"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	Database string `toml:"database"`
}

func (c *Database) MakeConnection() (string, string) {
	connection := fmt.Sprintf("%s:%s@/%s?parseTime=true", c.User, c.Pass, c.Database)
	return c.Driver, connection
}

var Databasesetting = &Database{}

func Setup() {
	tree, err := toml.LoadFile("conf/config.toml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf.config.toml': %v", err)
		return
	}

	serverTree := tree.Get("server").(*toml.Tree)
	serverTree.Unmarshal(Serversetting)

	dbTree := tree.Get("database").(*toml.Tree)
	dbTree.Unmarshal(Databasesetting)
}
