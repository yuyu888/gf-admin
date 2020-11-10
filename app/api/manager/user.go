package manager

import (
	"gf-admin/app/api"
	"gf-admin/app/service/manager"
	"gf-admin/app/service/sso"

	"github.com/gogf/gf/util/gconv"
)

type UserController struct {
	api.BaseController
}

func (c *UserController) List() {
	c.CheckLogin()
	mobile := c.Request.GetFormString("mobile")
	email := c.Request.GetFormString("email")
	real_name := c.Request.GetFormString("real_name")
	status := c.Request.GetFormInt("status", 0)

	limit := c.Request.GetFormInt("page_size", 20)
	page := c.Request.GetFormInt("page", 1)
	offset := (page - 1) * limit
	userService := new(manager.UserService)
	userService.SearchCond.Status = 1
	if len(mobile) > 0 {
		userService.SearchCond.Mobile = mobile
		userService.SearchCond.Status = 0
	}
	if len(email) > 0 {
		userService.SearchCond.Email = email
		userService.SearchCond.Status = 0
	}
	if len(real_name) > 0 {
		userService.SearchCond.RealName = real_name
		userService.SearchCond.Status = 0
	}
	if status > 0 {
		userService.SearchCond.Status = status
	}

	data, err := userService.List(limit, offset)
	if err != nil {
		c.Display(4001, "系统错误", nil)
	} else {
		c.Display(200, "success", data)
	}
}

func (c *UserController) Add() {
	mobile := c.Request.GetFormString("mobile")
	email := c.Request.GetFormString("email")
	real_name := c.Request.GetFormString("real_name")
	avatar := c.Request.GetFormString("avatar", "")
	password := c.Request.GetFormString("password", "123456")
	department := c.Request.GetFormString("department", "技术研发")
	roleids := c.Request.GetPost("roleids")

	if len(mobile) == 0 || len(real_name) == 0 || len(email) == 0 {
		c.Display(4001, "参数错误！", nil)
	}

	res, id := new(manager.UserService).Add(mobile, email, real_name, avatar, password, department)
	if res == true {
		var role_ids []int
		if roleids != nil {
			roleidMap := gconv.Map(roleids)
			for _, roleid := range roleidMap {
				role_ids = append(role_ids, gconv.Int(roleid))
			}
		}
		new(manager.RelationService).SetUserRole(gconv.Int(id), role_ids)

		c.Resp.Status = 200
		c.Resp.Message = "success"
		c.Resp.Data = id
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "添加失败"
	}
	c.Response.Write(c.Resp)
}

func (c *UserController) Edit() {
	uid := c.Request.GetFormInt("id", 0)
	mobile := c.Request.GetFormString("mobile", "")
	email := c.Request.GetFormString("email", "")
	avatar := c.Request.GetFormString("avatar", "")
	password := c.Request.GetFormString("password", "")
	department := c.Request.GetFormString("department", "")
	real_name := c.Request.GetFormString("real_name", "")
	roleids := c.Request.GetPost("roleids")
	if uid == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	var role_ids []int
	if roleids != nil {
		roleidMap := gconv.Map(roleids)
		for _, roleid := range roleidMap {
			role_ids = append(role_ids, gconv.Int(roleid))
		}
	}
	new(manager.RelationService).SetUserRole(uid, role_ids)
	res := new(manager.UserService).Edit(uid, mobile, email, avatar, password, department, real_name)
	if res == true {
		c.Resp.Status = 200
		c.Resp.Message = "success"
	} else {
		c.Resp.Status = 5001
		c.Resp.Message = "更新失败"
	}
	c.Response.Write(c.Resp)
}

func (c *UserController) Delete() {
	userid := c.Request.GetFormInt("id", 0)
	if userid == 0 {
		c.Display(4001, "参数错误！", nil)
	}
	res := new(manager.UserService).Delete(userid)
	if res == false {
		c.Display(5001, "删除失败", nil)
	}
	c.Display(200, "删除成功", nil)
}

func (c *UserController) Login() {
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

func (c *UserController) Loginout() {
	new(sso.SsoService).Loginout(c.Request)
	c.Display(200, "成功退出", nil)
}
