package manager

import (
	"github.com/gogf/gf/database/gdb"
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

func (relation *RelationModel) SaveRoleUser(roleid int, uid int) bool {
	if roleid == 0 || uid == 0 {
		return false
	}
	data := g.Map{"role_id": roleid, "uid": uid}
	_, err := g.DB().Table("admin_user_role_relation").Save(data)
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
	_, err := g.DB().Table("admin_role_menu_relation").Save(data)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) SaveRoleMenu(roleid int, menu_id int) bool {
	if roleid == 0 || menu_id == 0 {
		return false
	}
	data := g.Map{"role_id": roleid, "menu_id": menu_id}
	_, err := g.DB().Table("admin_role_menu_relation").Save(data)
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

func (relation *RelationModel) DeleteRoleUserExclude(uid int, roleids []int) bool {
	if uid == 0 {
		return false
	}
	_, err := g.DB().Table("admin_user_role_relation").Where("uid", uid).Where("role_id NOT IN(?)", roleids).Delete()
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

func (relation *RelationModel) DeleteRoleMenuExclude(menu_id int, roleids []int) bool {
	if menu_id == 0 {
		return false
	}
	_, err := g.DB().Table("admin_role_menu_relation").Where("menu_id", menu_id).Where("role_id NOT IN(?)", roleids).Delete()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (relation *RelationModel) MenuRole(menu_id int) (gdb.Result, error) {
	res, err := g.DB().Table("admin_role_menu_relation").Where("menu_id", menu_id).All()
	return res, err
}

func (relation *RelationModel) RoleMenu(role_id int) (gdb.Result, error) {
	res, err := g.DB().Table("admin_role_menu_relation").Where("role_id", role_id).All()
	return res, err
}

func (relation *RelationModel) UserMenus(uid int) (gdb.Result, error) {
	res, err := g.DB().Table("admin_role_menu_relation m").LeftJoin("admin_user_role_relation u", "m.role_id=u.role_id").Fields("m.*").Where("u.uid", uid).All()
	return res, err
}
