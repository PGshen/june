package web

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
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
	id, _ := strconv.Atoi(c.Param("id"))
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

}

// [API接口]更新API接口
func (apiWeb *SysApiWeb) EditApi(c *gin.Context) {

}

// [API接口]删除API接口
func (apiWeb *SysApiWeb) DelApi(c *gin.Context) {

}

// [API接口]分页获取API接口
func (apiWeb *SysApiWeb) ListApi(c *gin.Context) {

}

// [API接口]按ID获取API接口树
func (apiWeb *SysApiWeb) ApiTree(c *gin.Context) {

}

// [API接口]获取完整API接口树
func (apiWeb *SysApiWeb) ApiTreeById(c *gin.Context) {

}
