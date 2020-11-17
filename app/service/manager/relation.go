package manager

import (
	"gf-admin/app/model/manager"

	"github.com/gogf/gf/util/gconv"
)

type RelationService struct {
}

func (rs *RelationService) AddRoleUser(roleid int, uid int) bool {
	return new(manager.RelationModel).AddRoleUser(roleid, uid)
}

func (rs *RelationService) RoleMemberList(roleid int) (interface{}, error) {
	return new(manager.UserModel).RoleUserList(roleid)
}

func (rs *RelationService) DeleteRoleUser(roleid int, uid int) bool {
	return new(manager.RelationModel).DeleteRoleUser(roleid, uid)
}

func (rs *RelationService) SetUserRole(uid int, roleids []int) {
	new(manager.RelationModel).DeleteRoleUserExclude(uid, roleids)
	for _, roleid := range roleids {
		new(manager.RelationModel).SaveRoleUser(gconv.Int(roleid), uid)
	}
}

func (rs *RelationService) SetMenuRole(menu_id int, roleids []int) {
	new(manager.RelationModel).DeleteRoleMenuExclude(menu_id, roleids)
	for _, roleid := range roleids {
		new(manager.RelationModel).SaveRoleMenu(gconv.Int(roleid), menu_id)
	}
}

func (rs *RelationService) GetMenuRole(menu_id int) []int {
	res, err := new(manager.RelationModel).MenuRole(menu_id)
	var roleids []int
	if err == nil {
		for _, item := range res {
			roleids = append(roleids, gconv.Int(item["role_id"]))
		}
	}
	return roleids
}

func (rs *RelationService) GetRoleMenu(role_id int) []int {
	res, err := new(manager.RelationModel).RoleMenu(role_id)
	var menuids []int
	if err == nil {
		for _, item := range res {
			menuids = append(menuids, gconv.Int(item["menu_id"]))
		}
	}
	return menuids
}

func (rs *RelationService) AddRoleMenu(roleid int, menu_id int) bool {
	return new(manager.RelationModel).AddRoleMenu(roleid, menu_id)
}

func (rs *RelationService) DeleteRoleMenu(roleid int, menu_id int) bool {
	return new(manager.RelationModel).DeleteRoleMenu(roleid, menu_id)
}
