package manager

import (
	"gf-admin/app/api"

	"github.com/gogf/gf/frame/g"
)

type CommonController struct {
	api.BaseController
}

func (c *CommonController) Auth() {
	login_user := c.GetLoginUser()
	data := g.Map{"login_user": login_user}
	c.Display(200, "success", data)
}
