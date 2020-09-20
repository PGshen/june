package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysUser struct {
	UserId      int32     `gorm:"primary_key" json:"userId"`
	LoginName   string    `json:"loginName"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Avatar      string    `json:"avatar"`
	Remark      string    `json:"remark"`
	IsEnable    int       `json:"isEnable"`
	IsDel       int       `json:"isDel"`
	CreatedTime time.Time `json:"createdTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Roles       []int32   `sql:"-" json:"roles"`
}

func (SysUser) TableName() string {
	return "t_sys_user"
}

func (sysUser *SysUser) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now())
	err = scope.SetColumn("UpdateTime", time.Now())
	return err
}

func (sysUser *SysUser) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now())
	return err
}
