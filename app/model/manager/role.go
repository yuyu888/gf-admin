package manager

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type RoleModel struct {
}
type Role struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (r *RoleModel) Add(role_name string) (bool, int64) {
	role_info := g.Map{"name": role_name}
	res, err := g.DB().Table("admin_role").Insert(role_info)
	id, _ := res.LastInsertId()
	if err == nil {
		return true, id
	} else {
		return false, 0
	}
}

func (r *RoleModel) Edit(roleid int, role_name string) bool {
	update_data := g.Map{}

	if len(role_name) > 0 {
		update_data["name"] = role_name
	}

	if len(update_data) == 0 {
		return false
	}
	_, err := g.DB().Table("admin_role").Data(update_data).Where("id", roleid).Update()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (r *RoleModel) Delete(roleid int) bool {
	_, err := g.DB().Table("admin_role").Delete("id", roleid)
	if err == nil {
		g.DB().Table("admin_user_role_relation").Delete("role_id", roleid)
		g.DB().Table("admin_role_menu_relation").Delete("role_id", roleid)
		return true
	} else {
		return false
	}
}

func (r *RoleModel) List() (gdb.Result, error) {
	return g.DB().Table("admin_role").All()
}

func (r *RoleModel) UserRoleList(uid int) (gdb.Result, error) {
	res, err := g.DB().Table("admin_role r").LeftJoin("admin_user_role_relation ur", "r.id=ur.role_id").Fields("r.*").Where("ur.uid", uid).All()
	return res, err
}
