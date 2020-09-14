package web

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
	"github.com/PGshen/june/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysClientWeb struct {
	Log           logger.ILogger            `inject:""`
	ClientService service.ISysClientService `inject:""`
}

// [客户端]按ID获取客户端
func (clientWeb *SysClientWeb) GetClientById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		clientWeb.ClientService.GetClient(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			clientWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Client, ecode.P0308, errMsg, nil)
	}
}

// [客户端]新增客户端
func (clientWeb *SysClientWeb) SaveClient(c *gin.Context) {
	sysClient := models.SysClient{}
	err := c.BindJSON(&sysClient)
	// todo 参数校验
	if err != nil {
		clientWeb.Log.Error(err)
		resp.RespB406s(c, bcode.Client, ecode.P0308, "参数绑定错误", nil)
	} else {
		clientWeb.ClientService.SaveClient(c, &sysClient)
	}
}

// [客户端]更新客户端
func (clientWeb *SysClientWeb) EditClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysClient := models.SysClient{}
	err := c.BindJSON(&sysClient)
	sysClient.ClientId = int32(id)
	// todo 参数校验
	if err != nil {
		resp.RespB406s(c, bcode.Client, ecode.P0308, "参数绑定错误", nil)
	} else {
		clientWeb.ClientService.UpdateClient(c, &sysClient)
	}
}

// [客户端]删除客户端
func (clientWeb *SysClientWeb) DelClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		clientWeb.ClientService.DeleteClient(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			clientWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Client, ecode.P0308, errMsg, nil)
	}
}

// [客户端]分页获取客户端
func (clientWeb *SysClientWeb) ListClient(c *gin.Context) {
	reqCond := req.ReqCond{}
	err := c.BindJSON(&reqCond)
	if err != nil {
		resp.RespB406(c, bcode.Client, ecode.P0317, nil)
	} else {
		clientWeb.ClientService.ListClient(c, &reqCond)
	}
}

// [客户端]获取客户端绑定的IP
func (clientWeb *SysClientWeb) GetClientIp(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		clientWeb.ClientService.GetClientIp(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			clientWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Client, ecode.P0308, errMsg, nil)
	}
}

// [客户端]保存客户端绑定IP
func (clientWeb *SysClientWeb) SaveClientIp(c *gin.Context) {
	clientApiVo := vo.SysClientApiVo{}
	err := c.BindJSON(&clientApiVo)
	if err != nil {
		resp.RespB406(c, bcode.Client, ecode.P0317, nil)
	} else {
		clientWeb.ClientService.SaveClientIp(c, clientApiVo.ClientId, clientApiVo.Ip)
	}
}

// [客户端]删除客户端绑定IP
func (clientWeb *SysClientWeb) DelClientIp(c *gin.Context) {
	clientApiVo := vo.SysClientApiVo{}
	err := c.BindJSON(&clientApiVo)
	if err != nil {
		resp.RespB406(c, bcode.Client, ecode.P0317, nil)
	} else {
		clientWeb.ClientService.DelClientIp(c, clientApiVo.ClientId, clientApiVo.Ip)
	}
}

// [客户端]获取客户端关联的API
func (clientWeb *SysClientWeb) GetClientIpApi(c *gin.Context) {
	clientApiVo := vo.SysClientApiVo{}
	err := c.BindJSON(&clientApiVo)
	if err != nil {
		resp.RespB406(c, bcode.Client, ecode.P0317, nil)
	} else {
		clientWeb.ClientService.GetClientIpApi(c, clientApiVo.ClientId, clientApiVo.Ip)
	}
}

// [客户端]客户端IP关联的API(更新方式)
func (clientWeb *SysClientWeb) AuthClientIpApi(c *gin.Context) {
	clientApiVo := vo.SysClientApiVo{}
	err := c.BindJSON(&clientApiVo)
	if err != nil {
		resp.RespB406(c, bcode.Client, ecode.P0317, nil)
	} else {
		clientWeb.ClientService.AuthClientIpApi(c, clientApiVo.ClientId, clientApiVo.Ip, clientApiVo.AppId)
	}
}
