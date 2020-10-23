package task

import (
	"gf-admin/futurejob"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type FutureJob struct {
}

func (fj *FutureJob) Perform() {
	tl := NewTaskLock("FutureJobPerform")
	res := tl.Lock()
	if res == true {
		limit := 100
		dotask := true
		id := 0
		db := g.DB().Table("sys_future_job").Where("status", 1).Where("exec_time <=?", time.Now().Unix()).Limit(limit)
		for dotask {
			if id > 0 {
				db = db.Where("id>", id)
			}
			list, err := db.All()
			if err == nil && len(list) > 0 {
				var wg sync.WaitGroup
				for _, item := range list {
					wg.Add(1)
					id = gconv.Int(item["id"])
					go fj.evoke(id, gconv.String(item["job_code"]), gconv.String(item["params"]), &wg)
				}
				wg.Wait()
			}
			if len(list) < limit {
				dotask = false
			}
		}
	}
	tl.UnLock()
}

func (fj *FutureJob) evoke(id int, jobcode string, params string, wg *sync.WaitGroup) {
	defer wg.Done()
	futurejob.Scheduler(id, jobcode, params)
}
