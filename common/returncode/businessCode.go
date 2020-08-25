package returncode

type BusinessCode int32

// 长度两位
const (
	BusinessNormal   BusinessCode = 0
	BusinessAbnormal BusinessCode = 1
	User             BusinessCode = 2
	Role             BusinessCode = 3
	Menu             BusinessCode = 4
	Api              BusinessCode = 5
	Client           BusinessCode = 6
	Tenant           BusinessCode = 7
	Org              BusinessCode = 8
	Auth             BusinessCode = 9
	Token            BusinessCode = 10
	Permission       BusinessCode = 11
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
