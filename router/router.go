package router

import (
	"gf-admin/app/api/common"
	"gf-admin/app/api/manager"
	"gf-admin/app/service/sso"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	isLogin := new(sso.SsoService).CheckLogin(r)
	if isLogin == false {
		resp := `{"status":4030,"message":"没有登录"}`
		r.Response.WriteExit(resp)
	}
	r.Middleware.Next()
}

func init() {
	s := g.Server()

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareAuth)
		//group.ALL("/tools", new(tools.Controller), "Test, Localip")
		group.ALL("/manager/menu", new(manager.MenuController), "List, Add, Edit, Delete, FullList, MenuRole, RoleList")
		group.ALL("/manager/user", new(manager.UserController), "List, Add, Edit, Delete, FullList")
		group.ALL("/manager/role", new(manager.RoleController), "Add, Edit, Delete, List, AddRoleMember, MemberList, DeleteRoleMember, AddRoleMenu, DeleteRoleMenu")
	})
	s.BindController("/api/common/sso", new(common.SsoController), "Login, Loginout")
	s.BindController("/api/common/sys", new(common.SysController), "Auth")

}
