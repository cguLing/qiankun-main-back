package model

// Bearer 认证
type HttpHeader struct {
	Authorization   string `json:"Authorization" form:"Authorization"`
}

// 后端的响应
type HttpResponse struct {
	Code     int         `json:"code" form:"code"`
	Msg      string      `json:"msg" form:"msg"`
	Data     interface{} `json:"data" form:"data"`
}

