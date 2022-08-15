package model

/**
    model
    @author: roccoshi
    @desc: basic resp and req
**/

// Result Code
const (
	CodeOK  = 0
	CodeErr = -1
)

// BasicResp 基本返回值
type BasicResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewBasicResp 生成一个基本返回值
func NewBasicResp() *BasicResp {
	return &BasicResp{
		Code: CodeOK,
		Msg:  "success",
		Data: nil,
	}
}

// NewErrBasicResp 生成一个基本错误返回值
func NewErrBasicResp() *BasicResp {
	return &BasicResp{
		Code: CodeErr,
		Msg:  "error",
		Data: nil,
	}
}

type Conf struct {
	Redis RedisConf `yaml:"redis"`
}

type RedisConf struct {
	RequirePass string `yaml:"requirepass"`
}
