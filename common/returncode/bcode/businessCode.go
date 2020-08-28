package bcode

type BusinessCode int32

// 长度两位
const (
	BusinessNormal   BusinessCode = 0 * 1000
	BusinessAbnormal BusinessCode = 1 * 1000
	User             BusinessCode = 2 * 1000
	Role             BusinessCode = 3 * 1000
	Menu             BusinessCode = 4 * 1000
	Api              BusinessCode = 5 * 1000
	Client           BusinessCode = 6 * 1000
	Tenant           BusinessCode = 7 * 1000
	Org              BusinessCode = 8 * 1000
	Auth             BusinessCode = 9 * 1000
	Token            BusinessCode = 10 * 1000
	Permission       BusinessCode = 11 * 1000
)

var BusiCode = map[BusinessCode]string{
	BusinessNormal:   "正常",
	BusinessAbnormal: "异常",
	User:             "用户",
	Role:             "角色",
	Menu:             "菜单",
	Api:              "API",
	Client:           "客户端",
	Tenant:           "租户",
	Org:              "组织机构",
	Auth:             "认证",
	Token:            "令牌",
	Permission:       "权限",
}

func GetBusinessMsg(code BusinessCode) string {
	msg, ok := BusiCode[code]
	if ok {
		return msg
	}
	return BusiCode[BusinessNormal]
}
