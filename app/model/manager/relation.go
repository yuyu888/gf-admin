package manager

import (
	"github.com/gogf/gf/frame/g"
)

type RelationModel struct {
}

func (relation *RelationModel) AddRoleUser(roleid int, uid int) bool {
	if roleid == 0 || uid == 0 {
		return false
	}
	data := g.Map{"role_id": roleid, "uid": uid}
	_, err := g.DB().Table("admin_user_role_relation").Insert(data)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) AddRoleMenu(roleid int, menu_id int) bool {
	if roleid == 0 || menu_id == 0 {
		return false
	}
	data := g.Map{"role_id": roleid, "menu_id": menu_id}
	_, err := g.DB().Table("admin_role_menu_relation").Insert(data)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) DeleteRoleUser(roleid int, uid int) bool {
	if roleid == 0 || uid == 0 {
		return false
	}
	condition := g.Map{"role_id": roleid, "uid": uid}
	_, err := g.DB().Table("admin_user_role_relation").Where(condition).Delete()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) DeleteRoleUserAll(uid int) bool {
	if uid == 0 {
		return false
	}
	condition := g.Map{"uid": uid}
	_, err := g.DB().Table("admin_user_role_relation").Where(condition).Delete()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) DeleteRoleMenu(roleid int, menu_id int) bool {
	if roleid == 0 || menu_id == 0 {
		return false
	}
	condition := g.Map{"role_id": roleid, "menu_id": menu_id}
	_, err := g.DB().Table("admin_role_menu_relation").Where(condition).Delete()
	if err == nil {
		return true
	} else {
		return false
	}
}
