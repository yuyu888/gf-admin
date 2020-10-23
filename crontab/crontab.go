package crontab

import (
	"fmt"
	"gf-admin/crontab/task"
	"time"

	"github.com/gogf/gf/os/gcron"
)

func Register() {
	gcron.Add("@every 1m", new(task.FutureJob).Perform)
	gcron.Add("1 * * * * *", currentTime) //每分钟第一秒开始执行
}

func currentTime() {
	fmt.Println(time.Now())
}
