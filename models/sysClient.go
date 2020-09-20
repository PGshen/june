package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysClient struct {
	ClientId              int32  `gorm:"primary_key" json:"clientId"`
	AppId                 string `json:"appId"`
	AppSecret             string `json:"appSecret"`
	Description           string `json:"description"`
	ResourceIds           string `json:"resourceIds"`
	Scope                 string `json:"scope"`
	AuthorizedGrantTypes  string `json:"authorizedGrantTypes"`
	WebServerRedirectUri  string `json:"webServerRedirectUri"`
	Authorities           string `json:"authorities"`
	AccessTokenValidity   int32  `json:"accessTokenValidity"`
	RefreshTokenValidity  int32  `json:"refreshTokenValidity"`
	AdditionalInformation string `json:"additionalInformation"`
	Autoapprove           string `json:"autoapprove"`
	IsDel                 int    `json:"isDel"`

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
