package common

import (
	"gf-admin/app/api"
	"gf-admin/app/service/sso"
)

type SsoController struct {
	api.BaseController
}

func (c *SsoController) Login() {
	mobile := c.Request.GetFormString("mobile", "")
	password := c.Request.GetFormString("password", "")
	if len(mobile) == 0 || len(password) == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res, errstr := new(sso.SsoService).Login(mobile, password, c.Response)
	if res == true {
		c.Display(200, "成功登陆", nil)
	} else {
		c.Display(5001, errstr, nil)
	}
}

func (c *SsoController) Loginout() {
	new(sso.SsoService).Loginout(c.Request)
	c.Display(200, "成功退出", nil)
}
