package goft

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type TaskFunc func(params ...interface{})

// 异步任务列表
var TaskList chan *TaskExecutor
var once sync.Once

// 定时任务
var onceCron sync.Once
var taskCron *cron.Cron

func init() {
	chList := getTaskList()
	go func() {
		for t := range chList {
			doTask(t)
		}
	}()
}

// 任务执行器
type TaskExecutor struct {
	f        TaskFunc
	params   []interface{}
	callback func()
}

// 执行器
func (this *TaskExecutor) Exec() {
	this.f(this.params...)
}

// 开启协程进行处理。这里是web程序，后面看情况改成协程池
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

// 实力恶化一个构造器
func NewTaskExecutor(f TaskFunc, params []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{
		f:        f,
		params:   params,
		callback: callback,
	}
}

// 初始化任务列表
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		TaskList = make(chan *TaskExecutor, 0)
	})
	return TaskList
}

// 获取定时任务
func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

// 添加一个任务
func Task(f TaskFunc, callback func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, callback)
	}()
}
