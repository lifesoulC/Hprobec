package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func gorun(reqping ReqPing) (RespPing, error) {

	pool := new(GoroutinePool) //创建线程池
	pool.Init(10, len(reqping.ReqPairs))

	//	fmt.Println("Req in gorun")
	//	//fmt.Println(reqping.ReqPairs)
	//	fmt.Println("len is :=", len(reqping.ReqPairs))

	overtime := reqping.OverTime //发送超时
	fmt.Println("this is overtime")
	fmt.Println(overtime)

	for _, pairs := range reqping.ReqPairs {
		reqOpt := icmpReq{}
		reqOpt.Src = pairs.Source          //主机源地址
		reqOpt.Dest = pairs.Target         //目标地址
		reqOpt.Count = reqping.Count       //发送次数
		reqOpt.Interval = reqping.Interval //延迟时间
		reqOpt.TTL = reqping.TTL
		//		fmt.Println("this is req addtask")
		//		fmt.Println(reqOpt)
		b, err := json.Marshal(reqOpt)
		if err != nil {
			fmt.Println("json req error")
			break
		}
		pool.AddTask(func() error {
			return pool.httpDo(b)
		})

	}

	isFinish := false

	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})

	resp, err := pool.Start()
	if err != nil {
		pool.Stop()
		return resp, err
	}
	for !isFinish {
		time.Sleep(time.Millisecond * 100)
	}

	pool.Stop()

	fmt.Println("Over！")

	return resp, nil
}

func (self *GoroutinePool) httpDo(reqq []byte) error {

	req := icmpReq{}
	err := json.Unmarshal(reqq, &req)
	if err != nil {
		fmt.Println("json req error")
		return err

	}
	src := req.Src
	fmt.Println("this is req in http")
	fmt.Println(req)

	//dest := req.Dest
	//count := req.Count
	//time := req.Interval
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json req error")
		return err

	}
	src = "http://" + src + ":8080/probe/ping"
	client := &http.Client{}

	request, err := http.NewRequest("POST", src, bytes.NewReader(b))
	if err != nil {
		fmt.Println("NewRequset error")
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client error")
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io error")
		resp.Body.Close()
		return err
	}

	reqping := &icmpResp{}
	err = json.Unmarshal(body, reqping)
	if err != nil {
		fmt.Println("json error")
	}

	data := Data{}
	arr := reqping.Delays
	var sum = 0
	var max = 0
	var min = 99
	for _, vars := range arr {
		if max < vars {
			max = vars
		}
		if min > vars && vars != 0 {
			min = vars
		}
		sum = sum + vars
	}
	data.Max = max
	data.Min = min
	data.Avg = sum / reqping.Count
	data.Sum = sum
	data.Resp = *reqping

	self.result <- data //写入结果管道

	defer resp.Body.Close()

	return nil
}
