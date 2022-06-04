package utils

const (
	SUCCSE              = 0
	ERROR               = 1
	USER_REGISTER_FAILD = 20010003
)

var codeMsg = map[int]string{
	SUCCSE:              "Success",
	ERROR:               "Fail",
	USER_REGISTER_FAILD: "User Register Failed",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
