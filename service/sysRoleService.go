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

type ISysRoleService interface {
	GetRole(c *gin.Context, id int32)
	SaveRole(c *gin.Context, role *models.SysRole)
	UpdateRole(c *gin.Context, role *models.SysRole)
	DeleteRole(c *gin.Context, id int32)
	ListRole(c *gin.Context, reqCond *req.ReqCond)
	AllRole(c *gin.Context)
	GetRoleIdByUserId(c *gin.Context, userId int32)
	AuthRoleMenu(c *gin.Context, ra vo.RoleAuth)
}

type SysRoleService struct {
	Repo repository.ISysRoleRepo `inject:""`
	Log  logger.ILogger          `inject:""`
}

func (service *SysRoleService) GetRole(c *gin.Context, id int32) {
	service.Log.Infof("Get role: id = %s", id)
	data := service.Repo.GetRoleById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Role, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Role, data)
	}
}

func (service *SysRoleService) SaveRole(c *gin.Context, role *models.SysRole) {
	service.Log.Infof("Save role: role = %s", role)
	if service.Repo.InsertRole(role) {
		resp.RespB200(c, bcode.Role, role)
	} else {
		resp.RespB406(c, bcode.Role, ecode.P0313, nil)
	}
}

func (service *SysRoleService) UpdateRole(c *gin.Context, role *models.SysRole) {
	service.Log.Infof("Update role: role = %s", role)
	if service.Repo.UpdateRole(role) {
		resp.RespB200(c, bcode.Role, role)
	} else {
		resp.RespB406(c, bcode.Role, ecode.P0312, nil)
	}
}

func (service *SysRoleService) DeleteRole(c *gin.Context, id int32) {
	service.Log.Infof("Delete role: id = %s", id)
	// todo 判断是否有关联数据
	if service.Repo.DeleteRole(id) {
		resp.RespB200(c, bcode.Role, nil)
	} else {
		resp.RespB406(c, bcode.Role, ecode.P0302, nil)
	}
}

func (service *SysRoleService) ListRole(c *gin.Context, reqCond *req.ReqCond) {
	service.Log.Infof("List role, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := reqCond.Filter
	roles := service.Repo.ListRole(page, size, &total, where)
	resp.RespB200(c, bcode.Role, roles)
}

func (service *SysRoleService) AllRole(c *gin.Context) {
	roles := service.Repo.AllRole()
	resp.RespB200(c, bcode.Role, roles)
}

func (service *SysRoleService) GetRoleByUserId(userId int32) []*models.SysRole {
	return service.Repo.GetRoleByUserId(userId)
}

func (service *SysRoleService) GetRoleIdByUserId(c *gin.Context, userId int32) {
	data := service.Repo.GetRoleIdByUserId(userId)
	resp.RespB200(c, bcode.Role, data)
}

// 角色关联菜单
func (service *SysRoleService) AuthRoleMenu(c *gin.Context, ra vo.RoleAuth) {
	roleId := ra.RoleId
	menuIds := ra.MenuIds
	service.disassociateRoleMenu(roleId)
	service.associateRoleMenu(roleId, menuIds)
	resp.RespB200(c, bcode.Role, true)
}

func (service *SysRoleService) associateRoleMenu(roleId int32, menuIds []int32) {
	for e := range menuIds {
		service.Repo.SaveRoleMenu(roleId, int32(menuIds[e]))
	}
}

func (service *SysRoleService) disassociateRoleMenu(roleId int32) {
	service.Repo.DelRoleMenu(roleId)
}
