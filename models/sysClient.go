package models

import "time"

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
	TenantId    string    `json:"tenant_id"`
}

func (SysClient) TableName() string {
	return "t_sys_client"
}
