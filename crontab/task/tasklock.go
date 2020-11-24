package task

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
)

type TaskLock struct {
	prefix  string
	key     string
	expires int64
	value   string
}

func NewTaskLock(key string) *TaskLock {
	tl := &TaskLock{prefix: "TASKLOCK:"}
	tl.key = tl.prefix + key
	tl.expires = 600
	tl.value = grand.Letters(20)
	return tl
}

func (tl *TaskLock) Lock() bool {
	// N := grand.N(1, 400)
	// time.Sleep(time.Duration(N) * time.Millisecond)
	g.Redis().DoVar("SET", tl.key, tl.value, "NX", "EX", tl.expires)
	v, _ := g.Redis().DoVar("GET", tl.key)
	if v.String() != tl.value {
		return false
	}
	return true
}

func (tl *TaskLock) UnLock() {
	v, _ := g.Redis().DoVar("GET", tl.key)
	if v.String() == tl.value {
		g.Redis().Do("DEL", tl.key)
	}
}
