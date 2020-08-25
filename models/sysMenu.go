package models

import "time"

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
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}
