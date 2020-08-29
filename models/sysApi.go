package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysApi struct {
	ApiId       int32     `gorm:"primary_key" json:"api_id"`
	ParentApiId int32     `json:"parent_api_id"`
	CascadePath string    `json:"cascade_path"`
	Type        int       `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Method      string    `json:"method"`
	Uri         string    `json:"uri"`
	IsDel       int       `json:"is_del"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
	TenantId    string    `json:"tenant_id"`
}

func (SysApi) TableName() string {
	return "t_sys_api"
}

func (sysApi *SysApi) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now().Unix())
	return err
}

func (sysApi *SysApi) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now().Unix())
	return err
}
