package manager

import (
	"gf-admin/app/model/manager"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

type RoleService struct {
}

func (role *RoleService) List() (interface{}, error) {
	rolelist, err := new(manager.RoleModel).List()
	if err != nil {
		return nil, err
	}
	var data []interface{}
	for _, role := range rolelist {
		roleMap := gconv.Map(role)
		roleid := gconv.Int(roleMap["id"])
		members, err := new(manager.UserModel).RoleUserList(roleid)
		memberShow := ""
		if err == nil && len(members) > 0 {
			for _, m := range members {
				mMap := gconv.Map(m)
				memberShow = memberShow + gconv.String(mMap["real_name"]) + ","
			}
		}
		roleMap["members"] = memberShow
		data = append(data, roleMap)
	}
	return data, err
}

func (role *RoleService) Add(role_name string) (bool, int64) {
	return new(manager.RoleModel).Add(role_name)
}

func (role *RoleService) Delete(roleid int) bool {
	return new(manager.RoleModel).Delete(roleid)
}

func (role *RoleService) Edit(roleid int, role_name string) bool {
	return new(manager.RoleModel).Edit(roleid, role_name)
}

func (role *RoleService) RoleUserList(roleid int) (interface{}, error) {
	res, err := new(manager.UserModel).RoleUserList(roleid)
	return res, err
}

func (role *RoleService) OriginList() (gdb.Result, error) {
	return new(manager.RoleModel).List()
}
