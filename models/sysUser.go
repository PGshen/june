package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysUser struct {
	UserId      int32     `gorm:"primary_key" json:"user_id"`
	LoginName   string    `json:"login_name"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Avatar      string    `json:"avatar"`
	Remark      string    `json:"remark"`
	IsEnable    int       `json:"is_enable"`
	IsDel       int       `json:"is_del"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
	TenantId    string    `json:"tenant_id"`
}

func (SysUser) TableName() string {
	return "t_sys_user"
}

func (sysUser *SysUser) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now().Unix())
	return err
}

func (sysUser *SysUser) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now().Unix())
	return err
}
