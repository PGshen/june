package service

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysUserService interface {
	GetUserById(c *gin.Context, id int32)
	SaveUser(c *gin.Context, user *models.SysUser)
	UpdateUser(c *gin.Context, user *models.SysUser)
	DeleteUser(c *gin.Context, id int32)
	ListUser(c *gin.Context, reqCond *req.ReqCond)
	GetUserInfoByLoginName(loginName string) vo.UserInfoVo
	GetUserByLoginName(c *gin.Context, loginName string)
	CheckUser(loginName, password string) bool
	EnableUser(c *gin.Context, userId int32)
	AuthUserRole(c *gin.Context, ua vo.UserAuth)
}

type SysUserService struct {
	Repo     repository.ISysUserRepo `inject:""`
	RoleRepo repository.ISysRoleRepo `inject:""`
	ApiRepo  repository.ISysApiRepo  `inject:""`
	Log      logger.ILogger          `inject:""`
}

func (service *SysUserService) GetUserById(c *gin.Context, id int32) {
	service.Log.Infof("Get user: id = %s", id)
	data := service.Repo.GetUserById(id)
	roleIds := service.RoleRepo.GetRoleIdByUserId(id)
	data.Roles = roleIds
	if data == nil {
		resp.RespB406s(c, bcode.User, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.User, data)
	}
}

func (service *SysUserService) SaveUser(c *gin.Context, user *models.SysUser) {
	service.Log.Infof("Save user: user = %s", user)
	if service.Repo.InsertUser(user) {
		// 保存角色信息
		roleIds := user.Roles
		var userAuth = vo.UserAuth{
			UserId:  user.UserId,
			RoleIds: roleIds,
		}
		service.AuthUserRole(c, userAuth)
	} else {
		resp.RespB406(c, bcode.User, ecode.P0313, nil)
	}
}

func (service *SysUserService) UpdateUser(c *gin.Context, user *models.SysUser) {
	service.Log.Infof("Update user: user = %s", user)
	if service.Repo.UpdateUser(user) {
		// 保存角色信息
		roleIds := user.Roles
		var userAuth = vo.UserAuth{
			UserId:  user.UserId,
			RoleIds: roleIds,
		}
		service.AuthUserRole(c, userAuth)
	} else {
		resp.RespB406(c, bcode.User, ecode.P0312, nil)
	}
}

func (service *SysUserService) DeleteUser(c *gin.Context, id int32) {
	service.Log.Infof("Delete user: id = %s", id)
	// todo 判断是否有关联数据
	if service.Repo.DeleteUser(id) {
		resp.RespB200(c, bcode.User, nil)
	} else {
		resp.RespB406(c, bcode.User, ecode.P0302, nil)
	}
}

func (service *SysUserService) ListUser(c *gin.Context, reqCond *req.ReqCond) {
	service.Log.Infof("List user, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := reqCond.Filter
	users := service.Repo.ListUser(page, size, &total, where)
	res := make(map[string]interface{})
	res["records"] = users
	res["total"] = total
	resp.RespB200(c, bcode.User, res)
}

func (service *SysUserService) GetUserInfoByLoginName(loginName string) vo.UserInfoVo {
	userInfo := vo.UserInfoVo{}
	user := service.Repo.GetUserByLoginName(loginName)
	if user == nil {
		return userInfo
	}
	userInfo.User = *user
	// 角色
	roleIds := service.RoleRepo.GetRoleIdByUserId(user.UserId)
	userInfo.Roles = roleIds
	userInfo.User.Roles = roleIds
	// 权限
	apis := service.ApiRepo.GetApiByUserId(user.UserId)
	var permissions []string
	for e := range apis {
		api := apis[e]
		if api.Method != "" && api.Uri != "" {
			permissions = append(permissions, api.Method+":"+api.Uri)
		}
	}
	userInfo.Permissions = permissions
	return userInfo
}

// 根据登录名获取用户信息
func (service *SysUserService) GetUserByLoginName(c *gin.Context, loginName string) {
	userInfo := vo.UserInfoVo{}
	user := service.Repo.GetUserByLoginName(loginName)
	if user == nil {
		resp.RespB406(c, bcode.User, ecode.P0301, nil)
		return
	}
	userInfo.User = *user
	// 角色
	roleIds := service.RoleRepo.GetRoleIdByUserId(user.UserId)
	userInfo.Roles = roleIds
	userInfo.User.Roles = roleIds
	// 权限
	apis := service.ApiRepo.GetApiByUserId(user.UserId)
	var permissions []string
	for e := range apis {
		api := apis[e]
		if api.Method != "" && api.Uri != "" {
			permissions = append(permissions, api.Method+":"+api.Uri)
		}
	}
	userInfo.Permissions = permissions
	resp.RespB200(c, bcode.User, userInfo)
}

func (service *SysUserService) CheckUser(loginName, password string) bool {
	return service.Repo.CheckUser(loginName, password)
}

// 解禁/禁用用户
func (service *SysUserService) EnableUser(c *gin.Context, userId int32) {
	if service.Repo.EnableUser(userId) {
		resp.RespB200(c, bcode.User, true)
	} else {
		resp.RespB406(c, bcode.User, ecode.P0312, false)
	}
}

// 用户关联角色
func (service *SysUserService) AuthUserRole(c *gin.Context, ua vo.UserAuth) {
	userId := ua.UserId
	roleIds := ua.RoleIds
	service.disassociateUserRole(userId)
	service.associateUserRole(userId, roleIds)
	resp.RespB200(c, bcode.User, true)
}

func (service *SysUserService) associateUserRole(userId int32, roleIds []int32) {
	for e := range roleIds {
		service.Repo.SaveUserRole(userId, int32(roleIds[e]))
	}
}

func (service *SysUserService) disassociateUserRole(userId int32) {
	service.Repo.DelUserRole(userId)
}
