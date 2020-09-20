package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysMenu struct {
	MenuId       int32     `gorm:"primary_key" json:"menuId"`
	ParentMenuId int32     `json:"parentMenuId"`
	CascadePath  string    `json:"cascadePath"`
	MenuName     string    `json:"menuName"`
	Title        string    `json:"title"`
	Icon         string    `json:"icon"`
	Perm         string    `json:"perm"`
	Type         int       `json:"type"`
	OrderNum     int       `json:"orderNum"`
	Hidden       int       `json:"hidden"`
	AlwaysShow   int       `json:"alwaysShow"`
	Component    string    `json:"component"`
	Path         string    `json:"path"`
	Redirect     string    `json:"redirect"`
	IsDel        int       `json:"isDel"`
	CreatedTime  time.Time `json:"createdTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

type SysMenuTree struct {
	SysMenu
	Children []SysMenuTree `json:"children"`
}

func (SysMenu) TableName() string {
	return "t_sys_menu"
}

func (sysMenu *SysMenu) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now())
	err = scope.SetColumn("UpdateTime", time.Now())
	return err
}

func (sysMenu *SysMenu) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now())
	return err
}
