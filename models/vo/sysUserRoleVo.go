package vo

import "github.com/PGshen/june/models"

type UserAuth struct {
	UserId  int32   `json:"userId"`
	RoleIds []int32 `json:"roleIds"`
}

type UserInfoVo struct {
	User        models.SysUser `json:"user"`
	Permissions []string       `json:"permissions"`
	Roles       []int32        `json:"roles"`
}

type UserPermVo struct {
	UserId    int32    `json:"userId"`
	LoginName string   `json:"loginName"`
	Roles     []int32  `json:"roles"`
	Perms     []string `json:"perms"`
}
