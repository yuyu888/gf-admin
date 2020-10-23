package api

import "github.com/gogf/gf/frame/gmvc"

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
	c.Response.WriteExit(c.Resp)
}
