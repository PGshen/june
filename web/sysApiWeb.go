package web

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysApiWeb struct {
	Log        logger.ILogger         `inject:""`
	ApiService service.ISysApiService `inject:""`
}

// [API接口]按ID获取API接口
func (apiWeb *SysApiWeb) GetApiById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		apiWeb.ApiService.GetApi(c, id)
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			apiWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Api, ecode.P0308, errMsg, nil)
	}
}

// [API接口]新增API接口
func (apiWeb *SysApiWeb) SaveApi(c *gin.Context) {
	sysApi := models.SysApi{}
	err := c.BindJSON(&sysApi)
	// todo 参数校验
	if err != nil {
		apiWeb.Log.Error(err)
		resp.RespB406s(c, bcode.Api, ecode.P0308, "参数绑定错误", nil)
	} else {
		apiWeb.ApiService.SaveApi(c, &sysApi)
	}
}

// [API接口]更新API接口
func (apiWeb *SysApiWeb) EditApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysApi := models.SysApi{}
	err := c.BindJSON(&sysApi)
	sysApi.ApiId = int32(id)
	// todo 参数校验
	if err != nil {
		resp.RespB406s(c, bcode.Api, ecode.P0308, "参数绑定错误", nil)
	} else {
		apiWeb.ApiService.UpdateApi(c, &sysApi)
	}
}

// [API接口]删除API接口
func (apiWeb *SysApiWeb) DelApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		apiWeb.ApiService.DeleteApi(c, id)
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			apiWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Api, ecode.P0308, errMsg, nil)
	}
}

// [API接口]分页获取API接口
func (apiWeb *SysApiWeb) ListApi(c *gin.Context) {
	reqCond := req.ReqCond{}
	err := c.BindJSON(&reqCond)
	if err != nil {
		resp.RespB406(c, bcode.Api, ecode.P0317, nil)
	} else {
		apiWeb.ApiService.ListApi(c, &reqCond)
	}
}

// [API接口]获取完整API接口树
func (apiWeb *SysApiWeb) ApiTree(c *gin.Context) {
	apiWeb.ApiService.GetApiTrees(c)
}

// [API接口]按ID获取API接口树
func (apiWeb *SysApiWeb) ApiTreeById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		apiWeb.ApiService.GetApiTreeById(c, id)
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			apiWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Api, ecode.P0308, errMsg, nil)
	}
}
