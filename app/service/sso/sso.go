package sso

import (
	"fmt"
	"gf-admin/app/model/manager"
	"net/http"
	"time"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type SsoService struct {
}

type Authorization struct {
	LoginTime int64  `json:"login_time"`
	Signature string `json:"signature"`
	Uid       int    `json:"uid"`
	Expires   int    `json:"expires"`
}

func (s *SsoService) Login(mobile string, password string, HttpResp *ghttp.Response) (bool, string) {
	r, err := new(manager.UserModel).CheckUserByMp(mobile, password)
	if err == nil {
		if r == nil {
			return false, "用户名或者密码错误"
		} else {
			resMap := gconv.Map(r)
			uid := gconv.Int(resMap["id"])
			signature := buildSecretStr(gconv.String(uid), password)
			authorization := &Authorization{
				LoginTime: time.Now().Unix(),
				Signature: signature,
				Uid:       uid,
				Expires:   s.expiresTime(),
			}
			str := gconv.String(authorization)
			secret_key := s.secretKey()
			fmt.Print(str)

			_authorization, err := gaes.Encrypt(gconv.Bytes(str), gconv.Bytes(secret_key["key"]), gconv.Bytes(secret_key["iv"]))
			if err != nil {
				return false, "系统错误"
			}
			cookie_authorization := &http.Cookie{
				Name:     "Authorization",
				Value:    gbase64.EncodeToString(_authorization),
				MaxAge:   s.expiresTime(),
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(HttpResp.Writer, cookie_authorization)
			return true, "登录成功"
		}
	} else {
		return false, "系统错误"
	}
}

func (s *SsoService) CheckLogin(HttpReq *ghttp.Request) bool {
	authorization_str := HttpReq.Cookie.Get("Authorization")
	if len(authorization_str) == 0 {
		return false
	}
	authorization, err := gbase64.DecodeString(authorization_str)
	if err != nil {
		return false
	}
	secret_key := s.secretKey()
	_authorization, err := gaes.Decrypt(gconv.Bytes(authorization), gconv.Bytes(secret_key["key"]), gconv.Bytes(secret_key["iv"]))
	if err != nil {
		return false
	}
	auth := gjson.New(gconv.String(_authorization))
	if auth.Get("uid") == nil || auth.Get("signature") == nil {
		return false
	}
	res := lowCheck(gconv.String(auth.Get("uid")), gconv.String(auth.Get("signature")))
	return res
}

func (s *SsoService) Loginout(HttpReq *ghttp.Request) {
	HttpReq.Cookie.Remove("Authorization")
}

func (s *SsoService) CheckLoginHigh(HttpReq *ghttp.Request) bool {
	authorization_str := HttpReq.Cookie.Get("Authorization")
	if len(authorization_str) == 0 {
		return false
	}
	authorization, err := gbase64.DecodeString(authorization_str)
	if err != nil {
		return false
	}
	secret_key := s.secretKey()
	_authorization, err := gaes.Decrypt(gconv.Bytes(authorization), gconv.Bytes(secret_key["key"]), gconv.Bytes(secret_key["iv"]))
	if err != nil {
		return false
	}
	auth := gjson.New(gconv.String(_authorization))
	uid := gconv.String(auth.Get("uid"))
	signature := gconv.String(auth.Get("signature"))

	if len(uid) == 0 || len(signature) == 0 {
		return false
	}
	r, err := new(manager.UserModel).GetUserByUid(uid)
	if err != nil || r == nil {
		return false
	}
	resMap := gconv.Map(r)
	password := gconv.String(resMap["password"])
	res := highCheck(uid, password, signature)
	return res
}

func (s *SsoService) GetLoginUser(HttpReq *ghttp.Request) map[string]interface{} {
	authorization_str := HttpReq.Cookie.Get("Authorization")
	if len(authorization_str) == 0 {
		return g.Map{}
	}
	authorization, err := gbase64.DecodeString(authorization_str)
	if err != nil {
		return g.Map{}
	}
	secret_key := s.secretKey()
	_authorization, err := gaes.Decrypt(gconv.Bytes(authorization), gconv.Bytes(secret_key["key"]), gconv.Bytes(secret_key["iv"]))
	if err != nil {
		return g.Map{}
	}
	auth := gjson.New(gconv.String(_authorization))
	uid := gconv.String(auth.Get("uid"))
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

func (s *SsoService) expiresTime() int {
	return 3600
}

func (s *SsoService) secretKey() g.Map {
	return g.Map{
		"key": "c10157daa5b67bcd5bead831ba554e0c",
		"iv":  "yuyu888/gf-admin",
	}
}

func buildSecretStr(uid string, password string) string {
	lowstr := cMd5(uid)
	highstr := cMd5(uid + password)
	return lowstr + highstr
}

func cMd5(mystr string) string {
	secret_key := new(SsoService).secretKey()
	str, _ := gmd5.EncryptString(mystr + gconv.String(secret_key["key"]))
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
