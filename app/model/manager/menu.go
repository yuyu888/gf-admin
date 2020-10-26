package manager

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
)

type MenuModel struct {
}
type Menu struct {
	Id         int    `json:"id"`
	MenuName   string `json:"menu_name"`
	MenuPath   string `json:"menu_path"`
	MenuType   int    `json:"menu_type"`
	Fid        int    `json:"fid"`
	SortNo     int    `json:"sort_no"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (menu *MenuModel) Add(menu_name string, menu_path string, menu_type int, fid int, sort_no int) (bool, int) {
	_, err := g.DB().Table("admin_menu").Insert(g.Map{"menu_name": menu_name, "menu_path": menu_path, "menu_type": menu_type, "fid": fid, "sort_no": sort_no})
	if err == nil {
		return true, 1
	} else {
		return false, 0
	}
}

func (menu *MenuModel) Edit(menu_id int, menu_name string, menu_path string, menu_type int, sort_no int) bool {
	update_data := g.Map{}
	if len(menu_name) > 0 {
		update_data["menu_name"] = menu_name
	}
	if len(menu_path) > 0 {
		update_data["menu_path"] = menu_path
	}
	if menu_type > 0 {
		update_data["menu_type"] = menu_type
	}
	if sort_no > 0 {
		update_data["sort_no"] = sort_no
	}
	if len(update_data) == 0 {
		return false
	}
	_, err := g.DB().Table("admin_menu").Data(update_data).Where("id", menu_id).Update()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (menu *MenuModel) HasChildren(menu_id int) bool {
	r, err := g.DB().Table("admin_menu").Where("fid", menu_id).One()
	fmt.Println(r)

	if err == nil {
		if r != nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (menu *MenuModel) Delete(menu_id int) bool {
	r, err := g.DB().Table("admin_menu").Delete("id", menu_id)
	fmt.Println(r)

	if err == nil {
		return true
	} else {
		return false
	}
}
