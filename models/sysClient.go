package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysClient struct {
	ClientId              int32  `gorm:"primary_key" json:"client_id"`
	AppId                 string `json:"app_id"`
	AppSecret             string `json:"app_secret"`
	Description           string `json:"description"`
	ResourceIds           string `json:"resource_ids"`
	Scope                 string `json:"scope"`
	AuthorizedGrantTypes  string `json:"authorized_grant_types"`
	WebServerRedirectUri  string `json:"web_server_redirect_uri"`
	Authorities           string `json:"authorities"`
	AccessTokenValidity   int32  `json:"access_token_validity"`
	RefreshTokenValidity  int32  `json:"refresh_token_validity"`
	AdditionalInformation string `json:"additional_information"`
	Autoapprove           string `json:"autoapprove"`
	IsDel                 int    `json:"is_del"`

	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"modified_time"`
}

func (SysClient) TableName() string {
	return "t_sys_client"
}

func (sysClient *SysClient) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedTime", time.Now())
	err = scope.SetColumn("UpdateTime", time.Now())
	return err
}

func (sysClient *SysClient) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdateTime", time.Now())
	return err
}
