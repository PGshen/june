package service

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysMenuService interface {
	GetMenu(c *gin.Context, id int32)
	SaveMenu(c *gin.Context, menu *models.SysMenu)
	UpdateMenu(c *gin.Context, menu *models.SysMenu)
	DeleteMenu(c *gin.Context, id int32)
	ListMenu(c *gin.Context, reqCond *req.ReqCond)
	GetMenuByRoleId(c *gin.Context, roleId int32)
	GetMenuIdByRoleId(c *gin.Context, roleId int32)
	GetMenuTreeById(c *gin.Context, menuId int32)
	GetUserMenuTree(c *gin.Context, userId int32)
	GetUserMenuPerm(c *gin.Context, userId int32)
	GetMenuApiById(c *gin.Context, menuId int32)
	AuthMenuApi(c *gin.Context, ma vo.MenuAuth)
}

type SysMenuService struct {
	Repo       repository.ISysMenuRepo `inject:""`
	ApiService ISysApiService          `inject:""`
	Log        logger.ILogger          `inject:""`
}

func (service *SysMenuService) GetMenu(c *gin.Context, id int32) {
	service.Log.Infof("Get menu: id = %s", id)
	data := service.Repo.GetMenuById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Menu, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Menu, data)
	}
}

func (service *SysMenuService) SaveMenu(c *gin.Context, menu *models.SysMenu) {
	service.Log.Infof("Save menu: menu = %s", menu)
	if service.Repo.InsertMenu(menu) {
		resp.RespB200(c, bcode.Menu, menu)
	} else {
		resp.RespB406(c, bcode.Menu, ecode.P0313, nil)
	}
}

func (service *SysMenuService) UpdateMenu(c *gin.Context, menu *models.SysMenu) {
	service.Log.Infof("Update menu: menu = %s", menu)
	if service.Repo.UpdateMenu(menu) {
		resp.RespB200(c, bcode.Menu, menu)
	} else {
		resp.RespB406(c, bcode.Menu, ecode.P0312, nil)
	}
}

func (service *SysMenuService) DeleteMenu(c *gin.Context, id int32) {
	service.Log.Infof("Delete menu: id = %s", id)
	// todo 判断是否有关联数据
	if service.Repo.DeleteMenu(id) {
		resp.RespB200(c, bcode.Menu, nil)
	} else {
		resp.RespB406(c, bcode.Menu, ecode.P0302, nil)
	}
}

func (service *SysMenuService) ListMenu(c *gin.Context, reqCond *req.ReqCond) {
	service.Log.Infof("List menu, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := reqCond.Filter
	menus := service.Repo.ListMenu(page, size, &total, where)
	resp.RespB200(c, bcode.Menu, menus)
}

// 获取角色关联的菜单列表
func (service *SysMenuService) GetMenuByRoleId(c *gin.Context, roleId int32) {
	var menus []*models.SysMenu
	menus = service.Repo.GetMenuByRoleId(roleId)
	resp.RespB200(c, bcode.Menu, menus)
}

func (service *SysMenuService) GetMenuIdByRoleId(c *gin.Context, roleId int32) {
	var menuIds []int32
	menuIds = service.Repo.GetMenuIdByRoleId(roleId)
	resp.RespB200(c, bcode.Menu, menuIds)
}

// 获取菜单树
func (service *SysMenuService) GetMenuTreeById(c *gin.Context, menuId int32) {
	service.Log.Infof("Get menu tree by id, id = %s", menuId)
	menuTree := service.getMenuTree(menuId)
	if menuTree == nil {
		resp.RespB406(c, bcode.Menu, ecode.P0301, nil)
	} else {
		resp.RespB200(c, bcode.Menu, menuTree)
	}
}

// 递归查询，构建树结构
func (service *SysMenuService) getMenuTree(menuId int32) *models.SysMenuTree {
	var menuTree models.SysMenuTree
	menu := service.Repo.GetMenuById(menuId)
	if menu == nil {
		return nil
	}
	menuTree.SysMenu = *menu
	sysMenus := service.Repo.GetMenusByPid(menuId)
	// 递归查询子节点
	for child := range sysMenus {
		menuTree.Children = append(menuTree.Children, *service.getMenuTree(sysMenus[child].MenuId))
	}
	return &menuTree
}

// 指定用户的菜单树
func (service *SysMenuService) GetUserMenuTree(c *gin.Context, userId int32) {
	allMenu := service.getMenuTree(0).Children
	userMenuIds := service.Repo.GetMenuIdByUserId(userId)
	allMenu = service.cutOutMenuTree(allMenu, userMenuIds)
	resp.RespB200(c, bcode.Menu, allMenu)
}

// 裁剪菜单树，使其仅保留当前角色所关联的菜单
func (service *SysMenuService) cutOutMenuTree(menuTrees []models.SysMenuTree, menuIds []int32) []models.SysMenuTree {
	// 移除按钮类型
	for i := 0; i < len(menuTrees); {
		if menuTrees[i].Type == 2 {
			menuTrees = append(menuTrees[:i], menuTrees[i+1])
		} else {
			i++
		}
	}
	// 移除不在menuIds的
	for i := 0; i < len(menuTrees); {
		if !contains(menuIds, menuTrees[i].MenuId) {
			menuTrees = append(menuTrees[:i], menuTrees[i+1])
		} else {
			i++
		}
	}
	for i := 0; i < len(menuTrees); i++ {
		menuTrees[i].Children = service.cutOutMenuTree(menuTrees[i].Children, menuIds)
	}
	return menuTrees
}

func (service *SysMenuService) GetUserMenuPerm(c *gin.Context, userId int32) {
	var perms []string
	perms = service.Repo.GetMenuPermByUserId(userId)
	resp.RespB200(c, bcode.Menu, perms)
}

// 获取菜单关联的API
func (service *SysMenuService) GetMenuApiById(c *gin.Context, menuId int32) {
	var menuApiVo vo.SysMenuApiVo
	apiTree := service.ApiService.GetApiTree(1)
	// todo 深克隆，查询优化
	apiTree2 := service.ApiService.GetApiTree(1)
	//apiTree2 := &models.SysApiTree{}
	//_ = deepcopier.Copy(apiTree).To(apiTree2)
	apiTrees := []models.SysApiTree{*apiTree}
	apiTrees2 := []models.SysApiTree{*apiTree2}
	menuApiIds := service.Repo.GetApiIdByMenuId(menuId)
	fromData := service.cutApiTree(false, apiTrees, menuApiIds)
	toData := service.cutApiTree(true, apiTrees2, menuApiIds)
	menuApiVo.FromData = fromData
	menuApiVo.ToData = toData
	resp.RespB200(c, bcode.Menu, menuApiVo)
}

// 裁剪
func (service *SysMenuService) cutApiTree(flag bool, trees []models.SysApiTree, menuApiIds []int32) []models.SysApiTree {
	for e := range trees {
		if trees[e].Children != nil {
			trees[e].Children = service.cutApiTree(flag, trees[e].Children, menuApiIds)
		}
	}
	if flag {
		for i := 0; i < len(trees); {
			if (trees[i].Children == nil || len(trees[i].Children) == 0) && !contains(menuApiIds, trees[i].ApiId) {
				if i == len(trees)-1 {
					trees = trees[:i]
				} else {
					trees = append(trees[:i], trees[i+1])
				}
			} else {
				i++
			}
		}
	} else {
		for i := 0; i < len(trees); {
			if (trees[i].Children == nil || len(trees[i].Children) == 0) && contains(menuApiIds, trees[i].ApiId) {
				if i == len(trees)-1 {
					trees = trees[:i]
				} else {
					trees = append(trees[:i], trees[i+1])
				}
			} else {
				i++
			}
		}
	}
	return trees
}

// 绑定菜单关联的API
func (service *SysMenuService) AuthMenuApi(c *gin.Context, ma vo.MenuAuth) {
	menuId := ma.MenuId
	apiIds := ma.ApiIds
	service.disassociateMenuApi(menuId)
	service.associateMenuApi(menuId, apiIds)
	resp.RespB200(c, bcode.Menu, true)
}

func (service *SysMenuService) associateMenuApi(menuId int32, apiIds []int32) {
	for e := range apiIds {
		service.Repo.SaveMenuApi(menuId, int32(apiIds[e]))
	}
}

func (service *SysMenuService) disassociateMenuApi(menuId int32) {
	service.Repo.DelMenuApi(menuId)
}

func contains(slice []int32, ele int32) bool {
	for e := range slice {
		if slice[e] == ele {
			return true
		}
	}
	return false
}
