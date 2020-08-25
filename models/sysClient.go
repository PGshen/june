package models

import "time"

type SysClient struct {
	ClientId     int32     `gorm:"primary_key" json:"client_id"`
	AppId        string    `json:"app_id"`
	AppSecret    string    `json:"app_secret"`
	Description  string    `json:"description"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}
