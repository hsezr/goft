package goft

import (
	"sync"

	"github.com/robfig/cron/v3"
)


func init() {
	chlist := getTaskList() //得到任务列表
	go func() {
		for t := range chlist {
			doTask(t) //执行任务
		}
	} ()
}

func doTask(t *TaskExecutor) {
	go func ()  {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		} ()
		t.Exec()
	} ()
}

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once //单例模式
func getTaskList () chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskExecutor struct {
	f TaskFunc
	callback func()
	p []interface{} //参数
}

func (this *TaskExecutor) Exec() {
	this.f(this.p...)
}

func NewTaskExecutor(f TaskFunc, callback func(), p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}
func Task(f TaskFunc, cb func(),params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList()<-NewTaskExecutor(f, cb, params) //增加任务队列
	} ()
}

var onceCron sync.Once
var taskCron *cron.Cron //定时任务

func getCronTask() *cron.Cron {
	onceCron.Do(func () {
		taskCron = cron.New(cron.WithSeconds())
	})

	return taskCron
}

