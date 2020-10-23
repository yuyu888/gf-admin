package futurejob

import (
	"reflect"
)

var FuncMap = map[string]interface{}{
	"test": new(TestJob).Run,
}

func Scheduler(id int, jobcode string, params string) {
	funcValue := reflect.ValueOf(FuncMap[jobcode])
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(id), reflect.ValueOf(params)}
	// 反射调用函数
	funcValue.Call(paramList)
}
