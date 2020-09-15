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

type SysRoleWeb struct {
	Log         logger.ILogger          `inject:""`
	RoleService service.ISysRoleService `inject:""`
	MenuService service.ISysMenuService `inject:""`
}

// [角色]按ID获取角色
func (roleWeb *SysRoleWeb) GetRoleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		roleWeb.RoleService.GetRole(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			roleWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Role, ecode.P0308, errMsg, nil)
	}
}

// [角色]新增角色
func (roleWeb *SysRoleWeb) SaveRole(c *gin.Context) {
	sysRole := models.SysRole{}
	err := c.BindJSON(&sysRole)
	// todo 参数校验
	if err != nil {
		roleWeb.Log.Error(err)
		resp.RespB406s(c, bcode.Role, ecode.P0308, "参数绑定错误", nil)
	} else {
		roleWeb.RoleService.SaveRole(c, &sysRole)
	}
}

// [角色]更新角色
func (roleWeb *SysRoleWeb) EditRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysRole := models.SysRole{}
	err := c.BindJSON(&sysRole)
	sysRole.RoleId = int32(id)
	// todo 参数校验
	if err != nil {
		resp.RespB406s(c, bcode.Role, ecode.P0308, "参数绑定错误", nil)
	} else {
		roleWeb.RoleService.UpdateRole(c, &sysRole)
	}
}

// [角色]删除角色
func (roleWeb *SysRoleWeb) DelRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		roleWeb.RoleService.DeleteRole(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			roleWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Role, ecode.P0308, errMsg, nil)
	}
}

// [角色]分页获取角色
func (roleWeb *SysRoleWeb) ListRole(c *gin.Context) {
	reqCond := req.ReqCond{}
	err := c.BindJSON(&reqCond)
	if err != nil {
		resp.RespB406(c, bcode.Role, ecode.P0317, nil)
	} else {
		roleWeb.RoleService.ListRole(c, &reqCond)
	}
}

// [角色]获取全部角色
func (roleWeb *SysRoleWeb) GetAllRole(c *gin.Context) {
	roleWeb.RoleService.AllRole(c)
}

// [角色]按ID获取角色关联的菜单
func (roleWeb *SysRoleWeb) GetRoleMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		roleWeb.MenuService.GetMenuIdByRoleId(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			roleWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Role, ecode.P0308, errMsg, nil)
	}
}

// [角色]角色授权(角色关联菜单)
func (roleWeb *SysRoleWeb) AuthRoleMenu(c *gin.Context) {
	roleAuth := vo.RoleAuth{}
	err := c.BindJSON(&roleAuth)
	if err != nil {
		resp.RespB406s(c, bcode.Role, ecode.P0308, "参数绑定错误", nil)
	} else {
		roleWeb.RoleService.AuthRoleMenu(c, roleAuth)
	}
}
