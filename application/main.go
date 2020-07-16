package main

import (
	"log"
	"pdk/src/rest"
	"pdk/src/setting"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setting.Setup()
	log.Fatal(rest.RunAPI())
}
