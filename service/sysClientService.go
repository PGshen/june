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

type ISysClientService interface {
	GetClient(c *gin.Context, id int32)
	SaveClient(c *gin.Context, client *models.SysClient)
	UpdateClient(c *gin.Context, client *models.SysClient)
	DeleteClient(c *gin.Context, id int32)
	ListClient(c *gin.Context, reqCond *req.ReqCond)
}

type SysClientService struct {
	Repo repository.ISysClientRepo `inject:""`
	Log  logger.ILogger            `inject:""`
}

func (service *SysClientService) GetClient(c *gin.Context, id int32) {
	service.Log.Infof("Get client: id = %s", id)
	data := service.Repo.GetClientById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Client, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Client, data)
	}
}

func (service *SysClientService) SaveClient(c *gin.Context, client *models.SysClient) {
	service.Log.Infof("Save client: client = %s", client)
	if service.Repo.InsertClient(client) {
		resp.RespB200(c, bcode.Client, client)
	} else {
		resp.RespB406(c, bcode.Client, ecode.P0313, nil)
	}
}

func (service *SysClientService) UpdateClient(c *gin.Context, client *models.SysClient) {
	service.Log.Infof("Update client: client = %s", client)
	if service.Repo.UpdateClient(client) {
		resp.RespB200(c, bcode.Client, client)
	} else {
		resp.RespB406(c, bcode.Client, ecode.P0312, nil)
	}
}

func (service *SysClientService) DeleteClient(c *gin.Context, id int32) {
	service.Log.Infof("Delete client: id = %s", id)
	if service.Repo.DeleteClient(id) {
		resp.RespB200(c, bcode.Client, nil)
	} else {
		resp.RespB406(c, bcode.Client, ecode.P0302, nil)
	}
}

func (service *SysClientService) ListClient(c *gin.Context, reqCond *req.ReqCond) {
	service.Log.Infof("List client, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := reqCond.Filter
	clients := service.Repo.ListClient(page, size, &total, where)
	resp.RespB200(c, bcode.Client, clients)
}
