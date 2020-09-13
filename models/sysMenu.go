package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysMenu struct {
	MenuId       int32     `gorm:"primary_key" json:"menu_id"`
	ParentMenuId int32     `json:"parent_menu_id"`
	CascadePath  string    `json:"cascade_path"`
	MenuName     string    `json:"menu_name"`
	Title        string    `json:"title"`
	Icon         string    `json:"icon"`
	Perm         string    `json:"perm"`
	Type         int       `json:"type"`
	OrderNum     int       `json:"order_num"`
	Hidden       int       `json:"hidden"`
	AlwaysShow   int       `json:"always_show"`
	Component    string    `json:"component"`
	Path         string    `json:"path"`
	Redirect     string    `json:"redirect"`
	IsDel        int       `json:"is_del"`
	CreatedTime  time.Time `json:"created_time"`
	UpdateTime   time.Time `json:"update_time"`
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
