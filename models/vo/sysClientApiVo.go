package vo

import "github.com/PGshen/june/models"

type SysClientApiVo struct {
	ClientId int32   `json:"clientId"`
	Ip       string  `json:"ip"`
	ApiIds   []int32 `json:"apiIds"`
}

type SysApiVo struct {
	FromData []models.SysApiTree `json:"fromData"`
	ToData   []models.SysApiTree `json:"toData"`
}
