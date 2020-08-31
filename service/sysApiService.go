package service

import (
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysApiService interface {
	GetApi(c *gin.Context, id int)
	SaveApi(c *gin.Context, sysApi *models.SysApi)
	UpdateApi(c *gin.Context, sysApi *models.SysApi)
	DeleteApi(c *gin.Context, id int)
}

type SysApiService struct {
	Repo repository.ISysApiRepo `inject:""`
}

// 获取API
func (apiService *SysApiService) GetApi(c *gin.Context, id int) {
	data := apiService.Repo.GetApiById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Api, data)
	}
}

// 保存API
func (apiService *SysApiService) SaveApi(c *gin.Context, sysApi *models.SysApi) {
	if apiService.Repo.InsertApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0313, nil)
	}
}

// 更新API
func (apiService *SysApiService) UpdateApi(c *gin.Context, sysApi *models.SysApi) {
	if apiService.Repo.UpdateApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0312, nil)
	}
}

// 删除Api
func (apiService *SysApiService) DeleteApi(c *gin.Context, id int) {
	if apiService.Repo.DeleteApi(id) {
		resp.RespB200(c, bcode.Api, "")
	} else {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	}
}
