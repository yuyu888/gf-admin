package manager

import "gf-admin/app/model/manager"

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
	new(manager.RelationModel).DeleteRoleUserAll(uid)
	for _, roleid := range roleids {
		new(manager.RelationModel).AddRoleUser(roleid, uid)
	}
}
