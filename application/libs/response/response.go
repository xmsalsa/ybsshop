/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-07 10:23:43
 */
package response

type Response struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`
}

type MoreResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

var EmptyObjects = make(map[string]interface{})

func NewResponse(code int64, objects interface{}, msg string) *Response {
	return &Response{Code: code, Data: objects, Msg: msg}
}
func Resp() *Response {
	return &Response{Code: Success.Code, Data: EmptyObjects, Msg: Success.Msg}
}

func RespSuccess(objects interface{}) *Response {
	return &Response{Code: Success.Code, Data: objects, Msg: Success.Msg}
}
func RespFail(msg string) *Response {
	return &Response{Code: Fail.Code, Data: EmptyObjects, Msg: msg}
}
func RespValidatorFail(msg string) *Response {
	if msg == "" {
		msg = ValidatorFail.Msg
	}
	return &Response{Code: ValidatorFail.Code, Data: EmptyObjects, Msg: msg}
}
func RespDbExecutionFail(msg string) *Response {
	if msg == "" {
		msg = DbExecutionFail.Msg
	}
	msg = DbExecutionFail.Msg
	return &Response{Code: DbExecutionFail.Code, Data: EmptyObjects, Msg: msg}
}
func RespDbRecordNoExsit() *Response {
	return &Response{Code: DbRecordNoExsit.Code, Data: EmptyObjects, Msg: DbRecordNoExsit.Msg}
}

type ErrMsg struct {
	Code int64
	Msg  string
}

var (
	NoErr           = ErrMsg{2000, "请求成功"}
	AuthErr         = ErrMsg{4001, "认证错误"}
	AuthExpireErr   = ErrMsg{4002, "token 过期，请刷新token"}
	AuthActionErr   = ErrMsg{4003, "权限错误"}
	SystemErr       = ErrMsg{5000, "系统错误，请联系管理员"}
	DataEmptyErr    = ErrMsg{5001, "数据为空"}
	TokenCacheErr   = ErrMsg{5002, "TOKEN CACHE 错误"}
	Success         = ErrMsg{0, ""}
	Fail            = ErrMsg{1, "请求失败"}
	ValidatorFail   = ErrMsg{2, VALIDATOR_FAIL}
	DbExecutionFail = ErrMsg{4, DB_EXECUTION_FAIL}
	DbRecordNoExsit = ErrMsg{3, "不存在或已经删除"}
	DbRecordExsit   = ErrMsg{5, "已经存在"}
)
