package returncode

type ProgramCode int32

// 长度2位
var (
	ProgramNormal   ProgramCode = 0
	ProgramAbnormal ProgramCode = 1
	Db              ProgramCode = 2
	Queue           ProgramCode = 3
	Log             ProgramCode = 4
	Stream          ProgramCode = 5
)

var ProgCode = map[ProgramCode]string{
	ProgramNormal:   "正常",
	ProgramAbnormal: "异常",
	Db:              "数据库",
	Queue:           "队列",
	Log:             "日志",
	Stream:          "流",
}

func GetProgramMsg(code ProgramCode) string {
	msg, ok := ProgCode[code]
	if ok {
		return msg
	}
	return ProgCode[ProgramNormal]
}
