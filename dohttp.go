package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartHTTP(addr string) error {
	http.HandleFunc("/probe/ping", Doping)
	return http.ListenAndServe(addr, nil)
}

func Doping(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)

	reqping := ReqPing{}
	err := json.Unmarshal(body, &reqping)
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Println(reqping.Token)
	if reqping.Token != "T2PNzKP3T9oAwEka" {
		fmt.Println("token error")
		return
	}
	//	resp := &Resperr{}
	//	if err != nil {
	//		resp.ErrMsg = "json error"
	//		resp.ErrCode = errJson
	//		b, _ := json.Marshal(resp)
	//		w.Write(b)
	//		return
	//	}
	fmt.Println("this is reqping")
	fmt.Println(reqping)

	resp, err := gorun(reqping)
	if err != nil {
		fmt.Println("ERROR", err)
		//		resp.Code = 0
		//		//resp.ErrCode = 2
		//		resp.Data[0].Resp.ErrMsg = "client error"
		//		b, _ := json.Marshal(resp)
		s := "lcient error"
		b := []byte(s)
		w.Write(b)
		return

	}
	fmt.Println("this is resp return http")
	fmt.Println(resp)
	b, _ := json.Marshal(resp)
	w.Write(b)

}
