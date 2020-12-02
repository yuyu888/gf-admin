package futurejob

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type TestJob struct {
}

func (tj *TestJob) Run(id int, params string) {
	var data map[string]int
	fmt.Println(id)
	fmt.Println(params)

	json.Unmarshal([]byte(params), &data)
	fmt.Println(data)

	g.DB().Table("sys_future_job").Data(g.Map{"status": 2, "update_time": gtime.Now().Format("Y-m-d H:i:s")}).Where("id", id).Update()

}
