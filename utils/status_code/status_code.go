package status_code

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000  用户模块的错误
	ERROR_USERNAME_USED = 1001
)

var codeMsg = map[int]string{
	SUCCESS: "SUCCESS",
	ERROR:   "FAIL",

	ERROR_USERNAME_USED: "用户名已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
