package manager

import (
	"gf-admin/app/model/manager"
	"gf-admin/library/utils"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type MenuService struct {
}

type MenuTree struct {
	manager.Menu
	TypeShow string     `json:"type_show"`
	HasChild bool       `json:"has_child"`
	HasAuth  bool       `json:"has_auth"`
	Children []MenuTree `json:"children"`
}

var MenuTypeConfig = g.Map{"1": "菜单", "2": "接口", "3": "页面", "4": "按钮"}

func (rs *MenuService) MenuTree(fid int, mtype int) []MenuTree {
	var dataList []MenuTree

	menuList, err := new(manager.MenuModel).GetMenuByFid(fid, mtype)
	if err == nil && len(menuList) > 0 {
		for _, v := range menuList {
			var children []MenuTree
			mt := MenuTree{v, "", false, false, children}
			mt.TypeShow = gconv.String(MenuTypeConfig[gconv.String(mt.MenuType)])
			children = rs.MenuTree(v.Id, mtype)
			if len(children) > 0 {
				mt.HasChild = true
			}
			mt.Children = children

			dataList = append(dataList, mt)
		}
	}
	return dataList
}

func (menu *MenuService) Add(menu_name string, menu_path string, menu_type int, fid int, sort_no int) (bool, int64) {
	return new(manager.MenuModel).Add(menu_name, menu_path, menu_type, fid, sort_no)
}
func (menu *MenuService) Edit(menu_id int, menu_name string, menu_path string, menu_type int, sort_no int) bool {
	return new(manager.MenuModel).Edit(menu_id, menu_name, menu_path, menu_type, sort_no)
}

func (menu *MenuService) HasChildren(menu_id int) bool {
	return new(manager.MenuModel).HasChildren(menu_id)
}
func (menu *MenuService) Delete(menu_id int) bool {
	return new(manager.MenuModel).Delete(menu_id)
}

func (menu *MenuService) AuthMenuTree(fid int, mtype int, auth_menu_ids []int) []MenuTree {
	var dataList []MenuTree

	menuList, err := new(manager.MenuModel).GetMenuByFid(fid, mtype)
	if err == nil && len(menuList) > 0 {
		for _, v := range menuList {
			var children []MenuTree
			mt := MenuTree{v, "", false, false, children}
			mt.TypeShow = gconv.String(MenuTypeConfig[gconv.String(mt.MenuType)])
			children = menu.AuthMenuTree(v.Id, mtype, auth_menu_ids)
			if len(children) > 0 {
				mt.HasChild = true
			}
			mt.HasAuth = utils.InArray(mt.Id, auth_menu_ids)
			mt.Children = children

			dataList = append(dataList, mt)
		}
	}
	return dataList
}

func (menu *MenuService) UserMenuList(uid int) []MenuTree {
	var dataList []MenuTree
	auth_menu_ids := make([]int, 0)

	user_menus, err := new(manager.RelationModel).UserMenus(uid)
	if err == nil && len(user_menus) > 0 {
		for _, v := range user_menus {
			auth_menu_ids = append(auth_menu_ids, gconv.Int(v["menu_id"]))

		}
	}
	dataList = menu.AuthMenuTree(1, 1, auth_menu_ids)
	return dataList
}

func (menu *MenuService) CheckUserMenu(uid int, menu_path string) bool {
	auth_menu_ids := make([]int, 0)

	user_menus, err := new(manager.RelationModel).UserMenus(uid)
	if err == nil && len(user_menus) > 0 {
		for _, v := range user_menus {
			auth_menu_ids = append(auth_menu_ids, gconv.Int(v["menu_id"]))

		}
	}
	return true
}

func (menu *MenuService) UserMenuPath(uid int) map[string]bool {
	auth_menu_ids := make([]int, 0)

	user_menus, err := new(manager.RelationModel).UserMenus(uid)
	if err == nil && len(user_menus) > 0 {
		for _, v := range user_menus {
			auth_menu_ids = append(auth_menu_ids, gconv.Int(v["menu_id"]))

		}
	}
	dataList, err := new(manager.MenuModel).GetALL()
	paths := make(map[string]bool)
	for _, v := range dataList {
		if utils.InArray(gconv.Int(v["id"]), auth_menu_ids) == true {
			paths[gconv.String(v["menu_path"])] = true
		} else {
			paths[gconv.String(v["menu_path"])] = false
		}
	}
	return paths
}
