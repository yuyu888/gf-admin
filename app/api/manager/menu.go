package manager

import (
	"gf-admin/app/api"
)

type MenuController struct {
	api.BaseController
}

func (c *MenuController) List() {
	c.Display(4001, "something is wrong!", nil)
	c.Resp.Status = 200
	c.Resp.Message = "success"
	c.Response.Write(c.Resp)

	// c.Response.WriteExit("this is a test!!!")
}
