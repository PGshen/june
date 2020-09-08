package service

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
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
	ListApi(c *gin.Context, reqCond *req.ReqCond)
	GetApiTree(c *gin.Context)
	GetApiTreeById(c *gin.Context, id int)
}

type SysApiService struct {
	Repo repository.ISysApiRepo `inject:""`
	Log  logger.ILogger         `inject:""`
}

// 获取API
func (apiService *SysApiService) GetApi(c *gin.Context, id int) {
	apiService.Log.Info("Get api: id = %s", id)
	data := apiService.Repo.GetApiById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Api, data)
	}
}

// 保存API
func (apiService *SysApiService) SaveApi(c *gin.Context, sysApi *models.SysApi) {
	apiService.Log.Info("Save api: api = %s", sysApi)
	if apiService.Repo.InsertApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0313, nil)
	}
}

// 更新API
func (apiService *SysApiService) UpdateApi(c *gin.Context, sysApi *models.SysApi) {
	apiService.Log.Info("Update api: api = %s", sysApi)
	if apiService.Repo.UpdateApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0312, nil)
	}
}

// 删除Api
func (apiService *SysApiService) DeleteApi(c *gin.Context, id int) {
	apiService.Log.Info("Delete api: id = %s", id)
	if apiService.Repo.DeleteApi(id) {
		resp.RespB200(c, bcode.Api, "")
	} else {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	}
}

// 获取Api列表
func (apiService *SysApiService) ListApi(c *gin.Context, reqCond *req.ReqCond) {
	apiService.Log.Info("List api, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := reqCond.Filter
	apis := apiService.Repo.ListApi(page, size, &total, where)
	resp.RespB200(c, bcode.Api, apis)
}

// 完整API树
func (apiService *SysApiService) GetApiTree(c *gin.Context) {
	apiService.Log.Info("Get api tree")

}

// 指定ID的API树
func (apiService *SysApiService) GetApiTreeById(c *gin.Context, id int) {
	apiService.Log.Info("Get api tree by id, id = %s", id)
}
