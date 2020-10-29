package sso

import (
	"gf-admin/app/model/manager"
	"net/http"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type SsoService struct {
}

func (s *SsoService) Login(mobile string, password string, HttpResp *ghttp.Response) (bool, string) {
	r, err := new(manager.UserModel).CheckUserByMp(mobile, password)
	if err == nil {
		if r == nil {
			return false, "用户名或者密码错误"
		} else {
			resMap := gconv.Map(r)
			uid := gconv.String(resMap["id"])
			cookie_loginuid := &http.Cookie{
				Name:     "login_uid",
				Value:    uid,
				MaxAge:   36000,
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(HttpResp.Writer, cookie_loginuid)
			secret_key := buildSecretStr(uid, password)
			cookie_secretkey := &http.Cookie{
				Name:     "secret_key",
				Value:    secret_key,
				MaxAge:   36000,
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(HttpResp.Writer, cookie_secretkey)
			return true, "登录成功"
		}
	} else {
		return false, "系统错误"
	}
}

func (s *SsoService) CheckLogin(HttpReq *ghttp.Request) bool {
	uid := HttpReq.Cookie.Get("login_uid")
	secret_key := HttpReq.Cookie.Get("secret_key")
	if len(uid) == 0 || len(secret_key) == 0 {
		return false
	}
	res := lowCheck(uid, secret_key)
	return res
}

func (s *SsoService) Loginout(HttpReq *ghttp.Request) {
	HttpReq.Cookie.Remove("login_uid")
	HttpReq.Cookie.Remove("secret_key")
}

func (s *SsoService) CheckLoginHigh(HttpReq *ghttp.Request) bool {
	uid := HttpReq.Cookie.Get("login_uid")
	secret_key := HttpReq.Cookie.Get("secret_key")
	if len(uid) == 0 || len(secret_key) == 0 {
		return false
	}
	r, err := new(manager.UserModel).GetUserByUid(uid)
	if err != nil || r == nil {
		return false
	}
	resMap := gconv.Map(r)
	password := gconv.String(resMap["password"])
	res := highCheck(uid, password, secret_key)
	return res
}

func (s *SsoService) GetLoginUser(HttpReq *ghttp.Request) map[string]interface{} {
	uid := HttpReq.Cookie.Get("login_uid")
	r, err := new(manager.UserModel).GetUserByUid(uid)
	if err != nil || r == nil {
		return g.Map{}
	}
	resMap := gconv.Map(r)
	real_name := gconv.String(resMap["real_name"])
	avatar := gconv.String(resMap["avatar"])
	user_info := g.Map{"uid": uid, "real_name": real_name, "avatar": avatar}
	return user_info
}

func buildSecretStr(uid string, password string) string {
	lowstr := cMd5(uid)
	highstr := cMd5(uid + password)
	return lowstr + highstr
}

func cMd5(mystr string) string {
	key := "c10157daa5b67bcd5bead831ba554e0c"
	str, _ := gmd5.EncryptString(mystr + key)
	return str
}

func lowCheck(uid string, secretstr string) bool {
	str := cMd5(uid)
	if len(str) > len(secretstr) {
		return false
	}
	cstr := secretstr[0:len(str)]
	if str == cstr {
		return true
	} else {
		return false
	}
}

func highCheck(uid string, password string, secretstr string) bool {
	str := cMd5(uid + password)
	if len(str) > len(secretstr) {
		return false
	}
	cstr := secretstr[len(secretstr)-len(str) : len(secretstr)]
	if str == cstr {
		return true
	} else {
		return false
	}
}
