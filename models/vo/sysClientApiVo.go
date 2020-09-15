package vo

type SysClientApiVo struct {
	ClientId int32   `json:"clientId"`
	Ip       string  `json:"ip"`
	ApiIds   []int32 `json:"apiIds"`
}
