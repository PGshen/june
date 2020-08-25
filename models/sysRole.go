package models

import "time"

type SysRole struct {
	RoleId       int32     `gorm:"primary_key" json:"role_id"`
	RoleName     string    `json:"role_name"`
	Alias        string    `json:"alias"`
	Description  string    `json:"description"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}
