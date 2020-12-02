package boot

import (
	"flag"
	"fmt"
	"gf-admin/crontab"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
)

var (
	Env = flag.String("env", "prod", "set environment")
)

// 用于应用初始化。
func init() {
	// new(tools.Controller).Test()
	flag.Parse()
	runMode := genv.Get("RUN_MODE", "dev")

	fmt.Println(*Env)
	fmt.Println(runMode)
	// crontab.Register()
	if runMode == "prod" {
		crontab.Register() //生产环境下才运行定时任务
		g.SetDebug(false)
	}

	if runMode == "dev" {
		g.Cfg().SetFileName("config.dev.toml")
	}

	fmt.Println(g.Cfg().Get("test"))

	// s := g.Server()
	// // s.Plugin(&swagger.Swagger{})
}
