package utils

const (
	SUCCSE                   = 0
	ERROR                    = 1
	USERNAME_EMPTY           = 20010009
	PASSWORD_EMPTY           = 20010010
	USER_REGISTER_FAILD      = 20010003
	USER_NOT_REGISTER        = 20010008
	USER_LOGIN_FAILD         = 20010001
	USER_EDIT_PROFILE_FAILED = 20010005
	UPLOAD_PICTRUE_FAILD     = 30010005
	SAVE_PICTRUE_FAILD       = 30010006
)

var codeMsg = map[int]string{
	SUCCSE:                   "Success",
	ERROR:                    "Fail",
	USER_REGISTER_FAILD:      "User Register Failed",
	USER_LOGIN_FAILD:         "User Login Failed",
	USER_NOT_REGISTER:        "User Not Register",
	USERNAME_EMPTY:           "Username Empty",
	PASSWORD_EMPTY:           "Password Empty",
	USER_EDIT_PROFILE_FAILED: "User Edit Profile Failed",
	UPLOAD_PICTRUE_FAILD:     "Upload Picture Failed",
	SAVE_PICTRUE_FAILD:       "Save Picture Failed",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
