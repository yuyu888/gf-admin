package router

import (
	"gf-admin/app/api/manager"
	"gf-admin/app/api/tools"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.BindController("/api/tools", new(tools.Controller), "Test, Localip")
	s.BindController("/api/manager/menu", new(manager.MenuController), "List")

}
