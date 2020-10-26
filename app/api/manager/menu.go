package manager

import (
	"gf-admin/app/api"
	"gf-admin/app/model/manager"
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

func (c *MenuController) Add() {
	menu_name := c.Request.GetFormString("menu_name")
	fid := c.Request.GetFormInt("fid")
	menu_path := c.Request.GetFormString("menu_path")
	sort_no := c.Request.GetFormInt("sort_no", 1)
	menu_type := c.Request.GetFormInt("menu_type", 1)

	if len(menu_name) == 0 || fid == 0 || len(menu_path) == 0 || menu_type == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res, _ := new(manager.MenuModel).Add(menu_name, menu_path, menu_type, fid, sort_no)
	if res == true {
		c.Resp.Status = 200
		c.Resp.Message = "success"
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "添加失败"
	}
	c.Response.Write(c.Resp)
}

func (c *MenuController) Edit() {
	menu_id := c.Request.GetFormInt("id", 0)
	menu_name := c.Request.GetFormString("menu_name")
	menu_path := c.Request.GetFormString("menu_path")
	sort_no := c.Request.GetFormInt("sort_no", 0)
	menu_type := c.Request.GetFormInt("menu_type", 0)

	if len(menu_name) == 0 || menu_id == 0 || len(menu_path) == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res := new(manager.MenuModel).Edit(menu_id, menu_name, menu_path, menu_type, sort_no)
	if res == true {
		c.Resp.Status = 200
		c.Resp.Message = "success"
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "更新失败"
	}
	c.Response.Write(c.Resp)
}

func (c *MenuController) Delete() {
	menu_id := c.Request.GetFormInt("id", 0)
	if menu_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	hasChildren := new(manager.MenuModel).HasChildren(menu_id)
	if hasChildren == true {
		c.Display(4002, "请先删除子节点", nil)
	}
	res := new(manager.MenuModel).Delete(menu_id)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}

func (c *MenuController) FullList() {
	c.Display(4001, "something is wrong!", nil)
	c.Resp.Status = 200
	c.Resp.Message = "success"
	c.Response.Write(c.Resp)

	// c.Response.WriteExit("this is a test!!!")
}
