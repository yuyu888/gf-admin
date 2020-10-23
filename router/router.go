package router

import (
	"gf-admin/app/api/tools"

	"github.com/gogf/gf/frame/g"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()
	s.BindController("/api/tools", new(tools.Controller), "Test, Localip")
}
