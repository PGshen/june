package web

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysUserWeb struct {
	Log         logger.ILogger          `inject:""`
	UserService service.ISysUserService `inject:""`
	RoleService service.ISysRoleService `inject:""`
	MenuService service.ISysMenuService `inject:""`
}

// [用户]按ID获取用户
func (userWeb *SysUserWeb) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		userWeb.UserService.GetUserById(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			userWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.User, ecode.P0308, errMsg, nil)
	}
}

// [用户]新增用户
func (userWeb *SysUserWeb) SaveUser(c *gin.Context) {
	sysUser := models.SysUser{}
	err := c.BindJSON(&sysUser)
	// todo 参数校验
	if err != nil {
		userWeb.Log.Error(err)
		resp.RespB406s(c, bcode.User, ecode.P0308, "参数绑定错误", nil)
	} else {
		userWeb.UserService.SaveUser(c, &sysUser)
	}
}

// [用户]更新用户
func (userWeb *SysUserWeb) EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysUser := models.SysUser{}
	err := c.BindJSON(&sysUser)
	sysUser.UserId = int32(id)
	// todo 参数校验
	if err != nil {
		resp.RespB406s(c, bcode.User, ecode.P0308, "参数绑定错误", nil)
	} else {
		userWeb.UserService.UpdateUser(c, &sysUser)
	}
}

// [用户]删除用户
func (userWeb *SysUserWeb) DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		userWeb.UserService.DeleteUser(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			userWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.User, ecode.P0308, errMsg, nil)
	}
}

// [用户]分页获取用户
func (userWeb *SysUserWeb) ListUser(c *gin.Context) {
	reqCond := req.ReqCond{}
	err := c.BindJSON(&reqCond)
	if err != nil {
		resp.RespB406(c, bcode.User, ecode.P0317, nil)
	} else {
		userWeb.UserService.ListUser(c, &reqCond)
	}
}

// [用户]按ID获取用户角色
func (userWeb *SysUserWeb) GetUserRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userWeb.RoleService.GetRoleIdByUserId(c, int32(id))
}

// [用户]按ID禁用/解禁用户
func (userWeb *SysUserWeb) EnableUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userWeb.UserService.EnableUser(c, int32(id))
}

// [用户]修改用户密码
func (userWeb *SysUserWeb) ChangeUserPwd() {
	// todo
}

// [用户]按登录名获取用户
func (userWeb *SysUserWeb) GetUserByLoginName(c *gin.Context) {
	loginName := c.Param("loginName")
	userWeb.UserService.GetUserByLoginName(c, loginName)
}

// [登录]返回当前用户
func (userWeb *SysUserWeb) NowUser(c *gin.Context) {
	// todo
	userInfoVo := jwt.ExtractClaims(c)
	loginName := userInfoVo["loginName"].(string)
	userWeb.UserService.GetUserByLoginName(c, loginName)
}

// 获取用户的菜单树，用于生成前端菜单路由
func (userWeb *SysUserWeb) GetUserMenuTree(c *gin.Context) {
	userInfoVo := jwt.ExtractClaims(c)
	userId := int32(userInfoVo["userId"].(float64))
	userWeb.MenuService.GetUserMenuTree(c, userId)
}

// 获取用户的权限列表，用于前端控制权限按钮等
func (userWeb *SysUserWeb) GetUserPerm(c *gin.Context) {
	userInfoVo := jwt.ExtractClaims(c)
	userId := int32(userInfoVo["userId"].(float64))
	userWeb.MenuService.GetUserMenuPerm(c, userId)
}
