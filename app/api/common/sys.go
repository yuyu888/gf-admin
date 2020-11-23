package common

import (
	"gf-admin/app/api"
	"gf-admin/app/service/manager"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type SysController struct {
	api.BaseController
}

func (c *SysController) Auth() {
	c.CheckLoginHigh()
	login_user := c.GetLoginUser()
	menu_list := new(manager.MenuService).UserMenuList(gconv.Int(login_user["uid"]))
	menu_auth_list := new(manager.MenuService).UserMenuPath(gconv.Int(login_user["uid"]))
	data := g.Map{"login_user": login_user, "menu_list": menu_list, "menu_auth_list": menu_auth_list}
	c.Display(200, "success", data)
}
