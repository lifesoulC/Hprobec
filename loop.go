package main

import (
	"fmt"
	"sync"
)

type GoroutinePool struct {
	Queue          chan func() error
	Number         int
	Total          int
	mutex          *sync.Mutex
	result         chan Data
	finishCallback func()
}

// 初始化
func (self *GoroutinePool) Init(number int, total int) {
	self.Queue = make(chan func() error, total)
	self.Number = number
	self.Total = total
	self.mutex = &sync.Mutex{}
	self.result = make(chan Data, total)
}

func (self *GoroutinePool) Start() (RespPing, error) {
	// 开启Number个goroutine
	resp := RespPing{}

	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}

				err := task()
				if err != nil {
					fmt.Println("err")
					return
				}
			}
		}()
	}

	//获得每个work的执行结果
	for j := 0; j < self.Total; j++ {
		i, ok := <-self.result
		if !ok {
			break
		}
		resp.Code = 0
		resp.Data = append(resp.Data, i)

	}

	// 回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
	return resp, nil
}

func (self *GoroutinePool) Stop() {
	close(self.Queue)
	close(self.result)
}

// 添加任务
func (self *GoroutinePool) AddTask(task func() error) {

	self.Queue <- task
}

// 回调
func (self *GoroutinePool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}
