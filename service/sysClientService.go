package service

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/common/utils"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysClientService interface {
	GetClient(c *gin.Context, id int32)
	SaveClient(c *gin.Context, client *models.SysClient)
	UpdateClient(c *gin.Context, client *models.SysClient)
	DeleteClient(c *gin.Context, id int32)
	ListClient(c *gin.Context, reqCond *req.ReqCond)
	GetClientIp(c *gin.Context, id int32)
	SaveClientIp(c *gin.Context, clientId int32, ip string)
	DelClientIp(c *gin.Context, clientId int32, ip string)
	GetClientIpApi(c *gin.Context, clientId int32, ip string)
	AuthClientIpApi(c *gin.Context, clientId int32, ip string, apiIds []int32)
}

type SysClientService struct {
	Repo       repository.ISysClientRepo `inject:""`
	ApiRepo    repository.ISysApiRepo    `inject:""`
	ApiService ISysApiService            `inject:""`
	Log        logger.ILogger            `inject:""`
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
	where := utils.GetFilter(reqCond.Filter)
	clients := service.Repo.ListClient(page, size, &total, where)
	res := make(map[string]interface{})
	res["records"] = clients
	res["total"] = total
	resp.RespB200(c, bcode.Client, res)
}

func (service *SysClientService) GetClientIp(c *gin.Context, id int32) {
	service.Log.Infof("Get client's ip, clientId = %s", id)
	var ipsvo []vo.SysClientIpVo
	ips := service.Repo.GetClientIp(id)
	for e := range ips {
		ipsvo = append(ipsvo, vo.SysClientIpVo{Ip: ips[e]})
	}
	resp.RespB200(c, bcode.Client, ipsvo)
}

func (service *SysClientService) SaveClientIp(c *gin.Context, clientId int32, ip string) {
	// todo 检查IP是否已存在了
	// 绑定IP是默认填个apiId为0的以占位
	var apiIds []int32
	service.AuthClientIpApi(c, clientId, ip, apiIds)
}

func (service *SysClientService) DelClientIp(c *gin.Context, clientId int32, ip string) {
	if service.Repo.DelClientIp(clientId, ip) {
		resp.RespB200(c, bcode.Client, nil)
	} else {
		resp.RespB406s(c, bcode.Client, ecode.P0302, "取消关联IP失败", nil)
	}

}

func (service *SysClientService) GetClientIpApi(c *gin.Context, clientId int32, ip string) {
	apiIds := service.Repo.GetClientIpApi(clientId, ip)
	var apiVo vo.SysApiVo
	apiTree := service.ApiService.GetApiTree(1)
	// todo 深克隆，查询优化
	apiTree2 := service.ApiService.GetApiTree(1)
	//apiTree2 := &models.SysApiTree{}
	//_ = deepcopier.Copy(apiTree).To(apiTree2)
	apiTrees := []models.SysApiTree{*apiTree}
	apiTrees2 := []models.SysApiTree{*apiTree2}
	fromData := service.ApiService.CutApiTree(false, apiTrees, apiIds)
	toData := service.ApiService.CutApiTree(true, apiTrees2, apiIds)
	apiVo.FromData = fromData
	apiVo.ToData = toData
	resp.RespB200(c, bcode.Menu, apiVo)
}

func (service *SysClientService) AuthClientIpApi(c *gin.Context, clientId int32, ip string, apiIds []int32) {
	service.disassociateClientIpApi(clientId, ip)
	service.associateClientIpApi(clientId, ip, apiIds)
	resp.RespB200(c, bcode.Client, nil)
}

// 客户端关联API
func (service *SysClientService) associateClientIpApi(clientId int32, ip string, apiIds []int32) {
	apiIds = append(apiIds, 0)
	for e := range apiIds {
		service.Repo.SaveClientIpApi(clientId, ip, int32(apiIds[e]))
	}
}

// 客户端取消关联API
func (service *SysClientService) disassociateClientIpApi(clientId int32, ip string) {
	service.Repo.DelClientIp(clientId, ip)
}
