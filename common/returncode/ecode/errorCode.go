package ecode

type ErrorCode int32

// ErrorCode 总长度4位，前两位分类，后两位顺序
const (
	P0000 ErrorCode = 0

	P0101 ErrorCode = 101

	P0201 ErrorCode = 201
	P0202 ErrorCode = 202
	P0203 ErrorCode = 203
	P0204 ErrorCode = 204
	P0205 ErrorCode = 205

	P0301 ErrorCode = 301
	P0302 ErrorCode = 302
	P0303 ErrorCode = 303
	P0304 ErrorCode = 304
	P0305 ErrorCode = 305
	P0306 ErrorCode = 306
	P0307 ErrorCode = 307
	P0308 ErrorCode = 308
	P0309 ErrorCode = 309
	P0310 ErrorCode = 310
	P0311 ErrorCode = 311
	P0312 ErrorCode = 312
	P0313 ErrorCode = 313
	P0314 ErrorCode = 314
	P0315 ErrorCode = 315
	P0316 ErrorCode = 316
	P0317 ErrorCode = 317

	P0401 ErrorCode = 401
	P0402 ErrorCode = 402

	P0501 ErrorCode = 501
	P0502 ErrorCode = 502
	P0503 ErrorCode = 503
	P0504 ErrorCode = 504
	P0505 ErrorCode = 505
	P0506 ErrorCode = 506
	P0507 ErrorCode = 507
	P0508 ErrorCode = 508
	P0509 ErrorCode = 509
	P0510 ErrorCode = 510
	P0511 ErrorCode = 511
	P0512 ErrorCode = 512
	P0513 ErrorCode = 513
	P0514 ErrorCode = 514
	P0515 ErrorCode = 515
	P0516 ErrorCode = 516
	P0517 ErrorCode = 517
	P0518 ErrorCode = 518
	P0519 ErrorCode = 519
	P0520 ErrorCode = 520
	P0521 ErrorCode = 521
	P0522 ErrorCode = 522

	P0601 ErrorCode = 601
	P0602 ErrorCode = 602

	P0701 ErrorCode = 701
	P0702 ErrorCode = 702
)

var ErrCode = map[ErrorCode]string{
	/**
	 * 成功
	 */
	P0000: "成功",
	/**
	 * 系统错误 1开头
	 */
	P0101: "程序异常,请联系系统管理员",

	/**
	 * 文件类错误 2开头
	 */
	P0201: "文件不存在",
	P0202: "流操作异常",
	P0203: "文件格式校验失败",
	P0204: "文件内容校验失败",
	P0205: "文件大小超出限制",

	/**
	 * 数据类
	 */
	P0301: "数据不存在",
	P0302: "删除失败",
	P0303: "未更新",
	P0304: "有关联数据，不可删除",
	P0305: "数据深克隆错误",
	P0306: "数据已存在",
	P0307: "解析错误",
	P0308: "参数校验失败",
	P0309: "唯一性校验失败",
	P0310: "数据已存在",
	P0311: "部分记录失败",
	P0312: "数据更新失败",
	P0313: "数据写入失败",
	P0314: "数据删除失败",
	P0315: "远程查询返回失败",
	P0316: "数据解析失败",
	P0317: "数据绑定失败",

	/**
	 * 网络类
	 */
	P0401: "连接超时",
	P0402: "发送失败",

	/**
	 * 认证授权相关
	 */
	P0501: "用户名或密码错误",
	P0502: "令牌错误",
	P0503: "登录认证失败",
	P0504: "签名过期",
	P0505: "请求参数错误",
	P0506: "认证客户端信息异常",
	P0507: "拒绝访问",
	P0508: "未授权",
	P0509: "无效客户端",
	P0510: "未授权客户端",
	P0511: "无效的scope",
	P0512: "无效的token",
	P0513: "无效的请求",
	P0514: "重定向URL不匹配",
	P0515: "不支持的认证类型",
	P0516: "不支持的响应类型",
	P0517: "未知认证异常",
	P0518: "验证码错误",
	P0519: "账号被锁定，请联系管理员！",
	P0520: "操作过于频繁，请稍后再试！",
	P0521: "账号被禁用，请联系管理员！",
	P0522: "账号已过期，请联系管理员！",

	/**
	 * 服务类
	 */
	P0601: "服务不可用",
	P0602: "未找到对应方法",

	/**
	 * 批量/异步任务类
	 */
	P0701: "未找到任务",
	P0702: "任务调用失败",
}

func GetErrMsg(code ErrorCode) string {
	msg, ok := ErrCode[code]
	if ok {
		return msg
	}
	return ErrCode[P0000]
}

func GetErrorMsg(code int) string {
	errCode := ErrorCode(code)
	return GetErrMsg(errCode)
}
