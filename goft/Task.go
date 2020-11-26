package goft

import "sync"

type TaskFunc func(params ...interface{})

// 定义任务列表
var TaskList chan *TaskExecutor
var once sync.Once

func init() {
	chList := getTaskList()
	go func() {
		for t := range chList {
			t.Exec()
		}
	}()
}

// 任务执行器
type TaskExecutor struct {
	f      TaskFunc
	params []interface{}
}

// 执行器
func (this *TaskExecutor) Exec() {
	this.f(this.params...)
}

// 实力恶化一个构造器
func NewTaskExecutor(f TaskFunc, params []interface{}) *TaskExecutor {
	return &TaskExecutor{
		f:      f,
		params: params,
	}
}

// 初始化任务列表
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		TaskList = make(chan *TaskExecutor, 0)
	})
	return TaskList
}

// 添加一个任务
func Task(f TaskFunc, params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params)
	}()
}
