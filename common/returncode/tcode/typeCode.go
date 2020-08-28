package tcode

type TypeCode int32

//长度1位
var (
	TypeNormal TypeCode = 0 * 100000
	Program    TypeCode = 1 * 100000
	Business   TypeCode = 2 * 100000
)

var TypCode = map[TypeCode]string{
	TypeNormal: "正常",
	Program:    "程序",
	Business:   "业务",
}

func GetTypeMsg(code TypeCode) string {
	msg, ok := TypCode[code]
	if ok {
		return msg
	}
	return TypCode[TypeNormal]
}
