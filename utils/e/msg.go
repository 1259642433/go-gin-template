package e

var MsgFlags = map[int]string {
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	// 数据库
	ERROR_DATABASE: "数据库出错",

	// 其他
	ERROR_NO_FUC: "功能尚未开放",
	ERROR_UNKNOW: "未知错误",
	// token
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "用户名或密码错误",

	// banner
	ERROR_EXIST_BANNER:     "已存在该banner",
	ERROR_NOT_EXIST_BANNER: "该banner不存在",


}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}