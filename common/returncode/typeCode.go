package returncode

type TypeCode int32

//长度1位
var (
	TypeNormal TypeCode = 0
	Program    TypeCode = 1
	Business   TypeCode = 2
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
