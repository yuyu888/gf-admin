package api

import (
	"gf-admin/app/service/sso"

	"github.com/gogf/gf/frame/gmvc"
)

type BaseController struct {
	gmvc.Controller
	Resp ResopnseData
}

type ResopnseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *BaseController) Display(status int, message string, data interface{}) {
	c.Resp.Status = status
	c.Resp.Message = message
	c.Resp.Data = data
	c.Response.WriteExit(c.Resp)
}

func (c *BaseController) CheckLogin() {
	isLogin := new(sso.SsoService).CheckLogin(c.Request)
	if isLogin == false {
		c.Display(4030, "没有登录", nil)
	}
}

func (c *BaseController) CheckLoginHigh() {
	isLogin := new(sso.SsoService).CheckLoginHigh(c.Request)
	if isLogin == false {
		c.Display(4030, "没有登录", nil)
	}
}

func (c *BaseController) GetLoginUser() map[string]interface{} {
	c.CheckLogin()
	userinfo := new(sso.SsoService).GetLoginUser(c.Request)
	return userinfo
}
