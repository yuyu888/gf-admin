package manager

import (
	"gf-admin/app/api"
	"gf-admin/app/service/manager"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
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
	sort_no := c.Request.GetFormInt("sort_no", 0)
	menu_type := c.Request.GetFormInt("menu_type", 1)
	roleids := c.Request.GetPost("roleids")

	if len(menu_name) == 0 || fid == 0 || len(menu_path) == 0 || menu_type == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res, id := new(manager.MenuService).Add(menu_name, menu_path, menu_type, fid, sort_no)
	if res == true {
		var role_ids []int
		if roleids != nil {
			roleidMap := gconv.Map(roleids)
			for _, roleid := range roleidMap {
				role_ids = append(role_ids, gconv.Int(roleid))
			}
		}
		new(manager.RelationService).SetMenuRole(gconv.Int(id), role_ids)
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
	roleids := c.Request.GetPost("roleids")

	if len(menu_name) == 0 || menu_id == 0 || len(menu_path) == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	var role_ids []int
	if roleids != nil {
		roleidMap := gconv.Map(roleids)
		for _, roleid := range roleidMap {
			role_ids = append(role_ids, gconv.Int(roleid))
		}
	}
	new(manager.RelationService).SetMenuRole(menu_id, role_ids)
	res := new(manager.MenuService).Edit(menu_id, menu_name, menu_path, menu_type, sort_no)
	if res == true {
		c.Display(200, "更新成功", nil)
	} else {
		c.Display(5001, "更新失败", nil)
	}
}

func (c *MenuController) Delete() {
	menu_id := c.Request.GetFormInt("id", 0)
	if menu_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	hasChildren := new(manager.MenuService).HasChildren(menu_id)
	if hasChildren == true {
		c.Display(4002, "请先删除子节点", nil)
	}
	res := new(manager.MenuService).Delete(menu_id)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}

func (c *MenuController) FullList() {
	menu_list := new(manager.MenuService).MenuTree(1, 0)
	role_list, _ := new(manager.RoleService).OriginList()

	data := g.Map{"menu_list": menu_list, "role_list": role_list}
	c.Display(200, "success", data)
}

func (c *MenuController) MenuRole() {
	menu_id := c.Request.GetFormInt("id", 0)
	if menu_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	data := new(manager.RelationService).GetMenuRole(menu_id)
	c.Display(200, "success", data)
}

func (c *MenuController) RoleList() {
	role_id := c.Request.GetQueryInt("role_id", 0)
	if role_id == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	menu_ids := new(manager.RelationService).GetRoleMenu(role_id)

	menu_list := new(manager.MenuService).AuthMenuTree(1, 0, menu_ids)
	data := g.Map{"menu_list": menu_list}
	c.Display(200, "success", data)
}
