package models

import "time"

type SysApi struct {
	ApiId        int32     `gorm:"primary_key" json:"api_id"`
	ParentApiId  int32     `json:"parent_api_id"`
	CascadePath  string    `json:"cascade_path"`
	Type         int       `json:"type"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Method       string    `json:"method"`
	Uri          string    `json:"uri"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}

func (SysApi) TableName() string {
	return "t_sys_api"
}
