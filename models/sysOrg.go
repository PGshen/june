package models

import "time"

type SysOrg struct {
	OrgId       int32     `gorm:"primary_key" json:"orgId"`
	ParentOrgId int32     `json:"parentOrgId"`
	OrgName     string    `json:"orgName"`
	Remark      string    `json:"remark"`
	OrderNum    int       `json:"orderNum"`
	IsDel       int       `json:"isDel"`
	CreatedTime time.Time `json:"createdTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

func (SysOrg) TableName() string {
	return "t_sys_org"
}
