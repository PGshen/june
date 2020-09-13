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

type SysMenuWeb struct {
	Log         logger.ILogger          `inject:""`
	MenuService service.ISysMenuService `inject:""`
}

// [菜单]按ID获取菜单
func (menuWeb *SysMenuWeb) GetMenuById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		menuWeb.MenuService.GetMenu(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			menuWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Menu, ecode.P0308, errMsg, nil)
	}
}

// [菜单]新增菜单
func (menuWeb *SysMenuWeb) SaveMenu(c *gin.Context) {
	sysMenu := models.SysMenu{}
	err := c.BindJSON(&sysMenu)
	// todo 参数校验
	if err != nil {
		menuWeb.Log.Error(err)
		resp.RespB406s(c, bcode.Menu, ecode.P0308, "参数绑定错误", nil)
	} else {
		menuWeb.MenuService.SaveMenu(c, &sysMenu)
	}
}

// [菜单]更新菜单
func (menuWeb *SysMenuWeb) EditMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysMenu := models.SysMenu{}
	err := c.BindJSON(&sysMenu)
	sysMenu.MenuId = int32(id)
	// todo 参数校验
	if err != nil {
		resp.RespB406s(c, bcode.Menu, ecode.P0308, "参数绑定错误", nil)
	} else {
		menuWeb.MenuService.UpdateMenu(c, &sysMenu)
	}
}

// [菜单]删除菜单
func (menuWeb *SysMenuWeb) DelMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		menuWeb.MenuService.DeleteMenu(c, int32(id))
	} else {
		var errMsg = ""
		for _, err := range valid.Errors {
			menuWeb.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			errMsg += err.Message
		}
		resp.RespB406s(c, bcode.Menu, ecode.P0308, errMsg, nil)
	}
}

// [菜单]分页获取菜单
func (menuWeb *SysMenuWeb) ListMenu(c *gin.Context) {
	reqCond := req.ReqCond{}
	err := c.BindJSON(&reqCond)
	if err != nil {
		resp.RespB406(c, bcode.Menu, ecode.P0317, nil)
	} else {
		menuWeb.MenuService.ListMenu(c, &reqCond)
	}
}

// [菜单]按ID获取菜单树
func (menuWeb *SysMenuWeb) GetMenuTreeById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	menuWeb.MenuService.GetMenuTreeById(c, int32(id))
}

// [菜单]获取完整菜单树
func (menuWeb *SysMenuWeb) GetMenuTree(c *gin.Context) {
	menuWeb.MenuService.GetMenuTreeById(c, 0)
}

// [菜单]获取菜单关联的API
func (menuWeb *SysMenuWeb) GetMenuApiById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	menuWeb.MenuService.GetMenuApiById(c, int32(id))
}

// [菜单]菜单绑定API(更新方式)
func (menuWeb *SysMenuWeb) AuthMenuApi(c *gin.Context) {
	menuAuth := vo.MenuAuth{}
	err := c.BindJSON(&menuAuth)
	if err != nil {
		resp.RespB406s(c, bcode.Menu, ecode.P0308, "参数绑定错误", nil)
	} else {
		menuWeb.MenuService.AuthMenuApi(c, menuAuth)
	}
}
