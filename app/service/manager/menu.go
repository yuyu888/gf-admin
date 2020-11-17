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

func (rs *MenuService) AuthMenuTree(fid int, mtype int, auth_menu_ids []int) []MenuTree {
	var dataList []MenuTree

	menuList, err := new(manager.MenuModel).GetMenuByFid(fid, mtype)
	if err == nil && len(menuList) > 0 {
		for _, v := range menuList {
			var children []MenuTree
			mt := MenuTree{v, "", false, false, children}
			mt.TypeShow = gconv.String(MenuTypeConfig[gconv.String(mt.MenuType)])
			children = rs.AuthMenuTree(v.Id, mtype, auth_menu_ids)
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
