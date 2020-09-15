package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

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

func (sysRole *SysRole) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now())
	err = scope.SetColumn("UpdateTime", time.Now())
	return err
}

func (sysRole *SysRole) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now())
	return err
}
