package e

const (
	SUCCESS = 000000

	INVALID_PARAMS = 	100000

	URL_NOT_EXIST = 	100001

	ERROR_UNKNOWN = 	999999
)

var MsgFlag = map[int]string {
	SUCCESS: "ok",
	INVALID_PARAMS: "请求参数错误",
	URL_NOT_EXIST: "请求结果不存在",

	ERROR_UNKNOWN: "未知错误",
}

func GetErrorMessage(code int) string{
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ERROR_UNKNOWN]
}
