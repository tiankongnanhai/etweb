package utils

const (
	SUCCSE              = 0
	ERROR               = 1
	USERNAME_EMPTY      = 20010009
	PASSWORD_EMPTY      = 20010010
	USER_REGISTER_FAILD = 20010003
	USER_NOT_REGISTER   = 20010008
	USER_LOGIN_FAILD    = 20010001
)

var codeMsg = map[int]string{
	SUCCSE:              "Success",
	ERROR:               "Fail",
	USER_REGISTER_FAILD: "User Register Failed",
	USER_LOGIN_FAILD:    "User Login Failed",
	USER_NOT_REGISTER:   "User Not Register",
	USERNAME_EMPTY:      "Username Empty",
	PASSWORD_EMPTY:      "Password Empty",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
