package manager

import (
	"fmt"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
)

type UserModel struct {
}
type User struct {
	Id         int    `json:"id"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	RealName   string `json:"real_name"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	Status     int    `json:"status"`
	Department string `json:"department"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (user *UserModel) Add(mobile string, email string, real_name string, avatar string, password string, department string) (bool, int64) {
	md5_password, _ := gmd5.EncryptString(password)
	user_info := g.Map{"mobile": mobile, "email": email, "real_name": real_name, "avatar": avatar, "password": md5_password, "department": department}
	r, err := g.DB().Table("admin_user").Insert(user_info)
	id, _ := r.LastInsertId()
	if err == nil {
		fmt.Println(id)
		return true, id
	} else {
		return false, 0
	}
}

func (user *UserModel) Edit(uid int, mobile string, email string, avatar string, password string, department string, real_name string) bool {
	update_data := g.Map{}

	if len(mobile) > 0 {
		update_data["mobile"] = mobile
	}
	if len(email) > 0 {
		update_data["email"] = email
	}
	if len(real_name) > 0 {
		update_data["real_name"] = real_name
	}
	if len(avatar) > 0 {
		update_data["email"] = avatar
	}
	if len(password) > 0 {
		md5_password, _ := gmd5.EncryptString(password)
		update_data["password"] = md5_password
	}
	if len(avatar) > 0 {
		update_data["avatar"] = avatar
	}
	if len(department) > 0 {
		update_data["department"] = department
	}
	if len(update_data) == 0 {
		return false
	}
	_, err := g.DB().Table("admin_user").Data(update_data).Where("id", uid).Update()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (user *UserModel) Delete(uid int) bool {
	_, err := g.DB().Table("admin_user").Data(g.Map{"status": 2}).Where("id", uid).Update()
	if err == nil {
		return true
	} else {
		return false
	}
}
