package returncode

type HttpCode int32

// 长度3位，和标准HTTP状态码保持一致
const (
	Ok                 HttpCode = 200
	Created            HttpCode = 201
	Accepted           HttpCode = 202
	PartialContent     HttpCode = 206
	NotModified        HttpCode = 304
	Unauthorized       HttpCode = 401
	Forbidden          HttpCode = 403
	NotFound           HttpCode = 404
	NotAcceptable      HttpCode = 406
	Conflict           HttpCode = 409
	InternalError      HttpCode = 500
	ServiceUnavailable HttpCode = 503
)

var HttCode = map[HttpCode]string{
	Ok:                 "成功",
	Created:            "创建资源成功",
	Accepted:           "请求已接受",
	PartialContent:     "部分内容",
	NotModified:        "没有发生任何修改",
	Unauthorized:       "未授权",
	Forbidden:          "拒绝访问",
	NotFound:           "不存在所请求资源",
	NotAcceptable:      "请求未接受",
	Conflict:           "请求资源发生冲突",
	InternalError:      "服务器错误",
	ServiceUnavailable: "服务不可用",
}

// 由code获取错误信息
func GetHttpMsg(code HttpCode) string {
	msg, ok := HttCode[code]
	if ok {
		return msg
	}
	return HttCode[InternalError]
}
