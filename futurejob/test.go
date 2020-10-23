package futurejob

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/frame/g"
)

type TestJob struct {
}

func (tj *TestJob) Run(id int, params string) {
	var data map[string]int
	fmt.Println(id)
	fmt.Println(params)

	json.Unmarshal([]byte(params), &data)
	fmt.Println(data)
	t := time.Now()
	current_time := t.Format("2006-01-02 15:04:05")
	g.DB().Table("sys_future_job").Data(g.Map{"status": 2, "updated_time": current_time}).Where("id", id).Update()

}
