package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysUser struct {
	UserId       int32     `gorm:"primary_key" json:"user_id"`
	LoginName    string    `json:"login_name"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Avatar       string    `json:"avatar"`
	Remark       string    `json:"remark"`
	LockFlag     int       `json:"lock_flag"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}

func (sysUser *SysUser) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	return err
}

func (sysUser *SysUser) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	return err
}
