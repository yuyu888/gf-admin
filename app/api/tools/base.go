package tools

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/os/gtime"
)

type Controller struct {
	gmvc.Controller
}

type ResopnseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *Controller) Localip() {
	localIp, err := GetLocalIp()
	if err != nil {
		c.Response.Write("get local ip fail")
	}
	c.Response.Write(localIp)
}

func (c *Controller) ChangeEnv() {
	env := c.Request.GetQueryString("env")
	if env != "" {
		if env != "test" {
			env = "prod"
		}
		cookie := &http.Cookie{
			Name:   "env",
			Value:  env,
			MaxAge: 3600,
			Path:   "/",
		}

		http.SetCookie(c.Response.Writer, cookie)
	}
	var resp = ResopnseData{
		Status:  200,
		Message: "success",
	}
	c.Response.Write(resp)
}

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("get local ip error!")
}

func (c *Controller) Test() {
	db := g.DB().Table("sys_future_job").Where("status", 1).Where("exec_time <=?", time.Now().Unix()).Limit(100)
	list, _ := db.All()
	fmt.Println(list)
	ct := gtime.Now().Format("Y-m-d H:i:s")
	fmt.Println(ct)

	c.Response.WriteExit("this is a test!!!")
}
