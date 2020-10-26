package manager

import (
	"gf-admin/app/api"
	"gf-admin/app/model/manager"
)

type RoleController struct {
	api.BaseController
}

func (c *RoleController) List() {
	c.Display(4001, "something is wrong!", nil)
	c.Resp.Status = 200
	c.Resp.Message = "success"
	c.Response.Write(c.Resp)

	// c.Response.WriteExit("this is a test!!!")
}

func (c *RoleController) Add() {
	role_name := c.Request.GetFormString("role_name")

	if len(role_name) == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res, id := new(manager.RoleModel).Add(role_name)
	if res == true {
		c.Resp.Status = 200
		c.Resp.Message = "success"
		c.Resp.Data = id
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "添加失败"
	}
	c.Response.Write(c.Resp)
}

func (c *RoleController) Edit() {
	roleid := c.Request.GetFormInt("id", 0)
	role_name := c.Request.GetFormString("role_name", "")

	if roleid == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res := new(manager.RoleModel).Edit(roleid, role_name)
	if res == true {
		c.Resp.Status = 200
		c.Resp.Message = "success"
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "更新失败"
	}
	c.Response.Write(c.Resp)
}

func (c *RoleController) Delete() {
	roleid := c.Request.GetFormInt("id", 0)
	if roleid == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.RoleModel).Delete(roleid)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}

func (c *RoleController) AddRoleUser() {
	roleid := c.Request.GetFormInt("roleid", 0)
	uid := c.Request.GetFormInt("uid", 0)

	if roleid == 0 || uid == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.RelationModel).AddRoleUser(roleid, uid)
	if res == false {
		c.Display(5001, "添加失败", nil)
	}
	c.Display(200, "添加成功", nil)
}
func (c *RoleController) AddRoleMenu() {
	roleid := c.Request.GetFormInt("roleid", 0)
	menu_id := c.Request.GetFormInt("menu_id", 0)

	if roleid == 0 || menu_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.RelationModel).AddRoleMenu(roleid, menu_id)
	if res == false {
		c.Display(5001, "添加失败", nil)
	}
	c.Display(200, "添加成功", nil)
}
func (c *RoleController) DeleteRoleUser() {
	roleid := c.Request.GetFormInt("roleid", 0)
	uid := c.Request.GetFormInt("uid", 0)

	if roleid == 0 || uid == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.RelationModel).DeleteRoleUser(roleid, uid)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}
func (c *RoleController) DeleteRoleMenu() {
	roleid := c.Request.GetFormInt("roleid", 0)
	menu_id := c.Request.GetFormInt("menu_id", 0)

	if roleid == 0 || menu_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.RelationModel).DeleteRoleMenu(roleid, menu_id)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}
