package manager

import (
	"gf-admin/app/model/manager"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type UserService struct {
	SearchCond manager.UserSearchCond
}

func (u *UserService) List(limit int, offset int) (interface{}, error) {
	user_list, total, err := new(manager.UserModel).GetUserList(u.SearchCond, limit, offset)
	if err != nil {
		return nil, err
	}
	role_list, _ := new(manager.RoleModel).List()

	var userList []interface{}
	for _, user := range user_list {
		userMap := gconv.Map(user)
		uid := gconv.Int(userMap["id"])
		roles, err := new(manager.RoleModel).UserRoleList(uid)
		roleShow := ""
		var roleids []int
		if err == nil && len(roles) > 0 {
			for _, m := range roles {
				mMap := gconv.Map(m)
				roleShow = roleShow + gconv.String(mMap["name"]) + ","
				roleids = append(roleids, gconv.Int(mMap["id"]))
			}
		}
		userMap["roles"] = roleShow
		userMap["roleids"] = roleids
		userMap["role_list"] = roles
		userList = append(userList, userMap)
	}

	data := g.Map{"list": userList, "total": total, "role_list": role_list}
	return data, nil
}

func (u *UserService) Edit(uid int, mobile string, email string, avatar string, password string, department string, real_name string) bool {
	return new(manager.UserModel).Edit(uid, mobile, email, avatar, password, department, real_name)
}

func (u *UserService) Delete(uid int) bool {
	return new(manager.UserModel).Delete(uid)
}

func (u *UserService) Add(mobile string, email string, real_name string, avatar string, password string, department string) (bool, int64) {
	return new(manager.UserModel).Add(mobile, email, real_name, avatar, password, department)
}
