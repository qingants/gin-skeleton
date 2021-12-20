package errs

var ErrMsg = map[int]string{
	Success:                  "成功",
	Error:                    "失败",
	Maintenance:              "服务器维护",
	InvalidParams:            "参数错误",
}

func GetMsg(code int) string {
	msg, ok := ErrMsg[code]
	if ok {
		return msg
	}
	return ErrMsg[Error]
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetResponse(code int, data interface{}) Response {
	r := Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: data,
	}
	return r
}
