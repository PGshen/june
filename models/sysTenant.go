package models

import "time"

type SysTenant struct {
	TenantId    int32     `gorm:"primary_key" json:"tenant_id"`
	TenantName  string    `json:"tenant_name"`
	TenantCode  string    `json:"tenant_code"`
	Description string    `json:"description"`
	IsEnable    int       `json:"is_enable"`
	IsDel       int       `json:"is_del"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (SysTenant) TableName() string {
	return "t_sys_tenant"
}
