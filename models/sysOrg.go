package models

import "time"

type SysOrg struct {
	OrgId       int32     `gorm:"primary_key" json:"org_id"`
	ParentOrgId int32     `json:"parent_org_id"`
	OrgName     string    `json:"org_name"`
	Remark      string    `json:"remark"`
	OrderNum    int       `json:"order_num"`
	IsDel       int       `json:"is_del"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
	TenantId    string    `json:"tenant_id"`
}

func (SysOrg) TableName() string {
	return "t_sys_org"
}
