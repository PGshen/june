package vo

type SysClientApiVo struct {
	ClientId int32   `json:"clientId"`
	Ip       string  `json:"ip"`
	AppId    []int32 `json:"appId"`
}
