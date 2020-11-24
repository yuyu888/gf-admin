package router

import (
	"gf-admin/app/api/common"
	"gf-admin/app/api/manager"
	SM "gf-admin/app/service/manager"
	"gf-admin/app/service/sso"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func MiddlewareAuth(r *ghttp.Request) {
	isLogin := new(sso.SsoService).CheckLogin(r)
	if isLogin == false {
		resp := `{"status":4030,"message":"没有登录"}`
		r.Response.WriteExit(resp)
	}
	loginUser := new(sso.SsoService).GetLoginUser(r)
	menu_Auth_list := new(SM.MenuService).UserMenuPath(gconv.Int(loginUser["uid"]))
	url_path := strings.Split(r.RequestURI, "?")
	_, ok := menu_Auth_list[url_path[0]]
	if ok == true && menu_Auth_list[url_path[0]] == false {
		resp := `{"status":4031,"message":"没有权限"}`
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
	s.BindController("/api/common/sys", common.NewSysController(), "Auth")

}
