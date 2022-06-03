package utils

const (
	SUCCSE                   = 200
	ERROR                    = 500
	SERVER_ERROR             = 10000000
	INVALID_PARAMS           = 10000001
	NOT_FOUND                = 10000002
	USER_LOGIN_FALIED        = 20010001
	USER_LOGIN_REQUIRED      = 20010002
	USER_REGISTER_FALIED     = 20010003
	USER_GET_PROFILE_FAILED  = 20010004
	USER_EDIT_PROFILE_FAILED = 20010005
	UPLOAD_PICTURE_FAILED    = 30010001
)

var codeMsg = map[int]string{
	SUCCSE:                   "成功",
	ERROR:                    "失败",
	SERVER_ERROR:             "服务器错误",
	INVALID_PARAMS:           "不合法的请求参数",
	NOT_FOUND:                "没有发现",
	USER_LOGIN_FALIED:        "用户登陆失败",
	USER_LOGIN_REQUIRED:      "用户需要登陆",
	USER_REGISTER_FALIED:     "用户需要注册",
	USER_GET_PROFILE_FAILED:  "用户获取profile失败",
	USER_EDIT_PROFILE_FAILED: "用户编辑profile失败",
	UPLOAD_PICTURE_FAILED:    "上传图片失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
