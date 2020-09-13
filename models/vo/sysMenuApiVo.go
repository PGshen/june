package vo

import "github.com/PGshen/june/models"

type MenuAuth struct {
	MenuId int32   `json:"menuId"`
	ApiIds []int32 `json:"apiIds"`
}

type SysMenuApiVo struct {
	FromData []models.SysApiTree `json:"fromData"`
	ToData   []models.SysApiTree `json:"toData"`
}
