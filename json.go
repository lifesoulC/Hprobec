package main

type icmpReq struct {
	Token    string `json:"token"`
	Src      string `json:"src"`
	Dest     string `json:"dest"`
	Count    int    `json:"count"`
	Interval int    `json:"interval"`
	TTL      int    `json:"ttl"`
}

type icmpResp struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
	Token   string `json:"Token"`
	Src     string `json:"Src"`
	Dest    string `json:"Dest"`
	Delays  []int  `json:"Delays"`
	Count   int    `json:"Count"`
}

type Pairs struct {
	Server string `'json:"server"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type ReqPing struct {
	Token    string  `json:"token"`
	ReqPairs []Pairs `json:"reqpairs"`
	OverTime int     `json:"overtime"`
	Count    int     `json:"count"`
	Interval int     `json:"interval"`
	TTL      int     `json:"ttl"`
}

type Data struct {
	Min  int      `json:"min"`
	Avg  int      `json:"avg"`
	Max  int      `json:"max"`
	Sum  int      `json:"sum"`
	Resp icmpResp `json:"response"`
}

type RespPing struct {
	Code int `json:"code"`
	//Info string `json:"info"`
	Data []Data `json:"data"`
}
