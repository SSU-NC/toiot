package setting

import (
	"fmt"
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

type Health struct {
	Server string
}

func (hs *Health) Getenv() {
	hs.Server = os.Getenv("HEALTH_SERVER")
	if hs.Server == "" {
		hs.Server = "0.0.0.0:8083"
	}
}

var Healthsetting = &Health{}

type App struct {
	Server      string
	RequestPath string
}

func (as *App) Getenv() {
	as.Server = os.Getenv("APP_SERVER")
	if as.Server == "" {
		as.Server = "localhost:8081"
	}
	as.RequestPath = "/regist/sink"
}

var Appsetting = &App{}

type Status struct {
	Count int
	Tick  int
	Drop  int
}

func (ss *Status) Getenv() {
	GetenvInt(&ss.Count, 5, "STATUS_COUNT")
	GetenvInt(&ss.Tick, 60, "STATUS_TICK")
	GetenvInt(&ss.Drop, 1, "STATUS_DROP")
}

var StatusSetting = &Status{}

func init() {
	Healthsetting.Getenv()
	Appsetting.Getenv()
	StatusSetting.Getenv()

	fmt.Printf("Health : &v\nApp : %v\nStatus : %v\n\n", Healthsetting, Appsetting, StatusSetting)
}
