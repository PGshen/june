package models

import "time"

type SysRole struct {
	RoleId      int32     `gorm:"primary_key" json:"role_id"`
	RoleName    string    `json:"role_name"`
	Alias       string    `json:"alias"`
	Description string    `json:"description"`
	IsDel       int       `json:"is_del"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (SysRole) TableName() string {
	return "t_sys_role"
}
