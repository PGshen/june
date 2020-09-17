package resp

import (
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/common/returncode/hcode"
	"github.com/PGshen/june/common/returncode/pcode"
	"github.com/PGshen/june/common/returncode/tcode"
	"github.com/PGshen/june/common/setting"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ResponseData 数据返回结构体
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/**
 * 406
 */
func RespB406(c *gin.Context, businessCode bcode.BusinessCode, errorCode ecode.ErrorCode, data interface{}) {
	RespBusiData(c, hcode.NotAcceptable, tcode.Business, businessCode, errorCode, "", data)
}

func RespB406s(c *gin.Context, businessCode bcode.BusinessCode, errorCode ecode.ErrorCode, errorMsg string, data interface{}) {
	RespBusiData(c, hcode.NotAcceptable, tcode.Business, businessCode, errorCode, errorMsg, data)
}

func RespP406(c *gin.Context, programCode pcode.ProgramCode, errorCode ecode.ErrorCode, data interface{}) {
	RespProgramData(c, hcode.NotAcceptable, tcode.Business, programCode, errorCode, "", data)
}

func RespP406s(c *gin.Context, programCode pcode.ProgramCode, errorCode ecode.ErrorCode, errorMsg string, data interface{}) {
	RespProgramData(c, hcode.NotAcceptable, tcode.Business, programCode, errorCode, errorMsg, data)
}

/**
 * 500
 */
func RespP500(c *gin.Context, programCode pcode.ProgramCode, errorCode ecode.ErrorCode, data interface{}) {
	RespProgramData(c, hcode.InternalError, tcode.Business, programCode, errorCode, "", data)
}

/**
 * 200
 */
func RespB200(c *gin.Context, businessCode bcode.BusinessCode, data interface{}) {
	RespBusiData(c, hcode.Ok, tcode.Business, businessCode, ecode.P0000, "", data)
}
func RespB200m(c *gin.Context, businessCode bcode.BusinessCode, msg string, data interface{}) {
	RespBusiData(c, hcode.Ok, tcode.Business, businessCode, ecode.P0000, msg, data)
}

func Resp200(c *gin.Context, businessCode bcode.BusinessCode) {
	RespBusiData(c, hcode.Ok, tcode.Business, businessCode, ecode.P0000, "", nil)
}

// 业务逻辑错误
func RespBusiData(c *gin.Context, httpCode hcode.HttpCode, typeCode tcode.TypeCode, businessCode bcode.BusinessCode, errorCode ecode.ErrorCode, errorMsg string, data interface{}) {
	businessMsg := bcode.GetBusinessMsg(businessCode)
	if errorMsg == "" {
		errorMsg = ecode.GetErrMsg(errorCode)
	}
	resp := ResponseData{
		Code:    int(httpCode) + int(typeCode) + int(businessCode) + int(errorCode),
		Message: businessMsg + errorMsg,
		Data:    data,
	}
	//RespJSON(c, int(httpCode)/1000000, resp)
	RespJSON(c, 200, resp)
}

// 程序错误
func RespProgramData(c *gin.Context, httpCode hcode.HttpCode, typeCode tcode.TypeCode, programCode pcode.ProgramCode, errorCode ecode.ErrorCode, errorMsg string, data interface{}) {
	businessMsg := pcode.GetProgramMsg(programCode)
	if errorMsg == "" {
		errorMsg = ecode.GetErrMsg(errorCode)
	}
	resp := ResponseData{
		Code:    int(httpCode) + int(typeCode) + int(programCode) + int(errorCode),
		Message: businessMsg + errorMsg,
		Data:    data,
	}
	//RespJSON(c, int(httpCode)/1000000, resp)
	RespJSON(c, 200, resp)
}

//RespJSON 返回JSON数据
func RespJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

//GetPage 获取每页数量
func GetPage(c *gin.Context) (page, pagesize int) {
	page, _ = strconv.Atoi(c.Query("page"))
	pagesize, _ = strconv.Atoi(c.Query("limit"))
	if pagesize == 0 {
		pagesize = setting.Config.App.PageSize
	}
	if page == 0 {
		page = 1
	}
	return
}
