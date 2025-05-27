package R

/*统一返回*/
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	SUCCESS_CODE = 200
	FAIL_CODE    = 40001
	SUCCESS_MSG  = "success"
	FAIL_MSG     = "fail"
)

func Ok(data any) *Response {
	return &Response{Code: SUCCESS_CODE, Message: SUCCESS_MSG, Data: data}
}

func OkCode(code int, data any) *Response {
	return &Response{Code: code, Message: SUCCESS_MSG, Data: data}
}

func OkCodeMsg(code int, msg string, data any) *Response {
	return &Response{Code: SUCCESS_CODE, Message: msg, Data: data}
}

func Fail() *Response {
	return &Response{Code: FAIL_CODE, Message: FAIL_MSG, Data: nil}
}

func FailCodeMsg(code int, msg string) *Response {
	return &Response{Code: code, Message: msg, Data: nil}
}
