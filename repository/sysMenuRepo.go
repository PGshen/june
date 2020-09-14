package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
)

type ISysMenuRepo interface {
	GetMenuById(id int32) *models.SysMenu
	InsertMenu(menu *models.SysMenu) bool
	UpdateMenu(menu *models.SysMenu) bool
	DeleteMenu(id int32) bool
	ListMenu(page, size int32, total *int32, where interface{}) []*models.SysMenu
	GetMenuByRoleId(roleId int32) []*models.SysMenu
	GetMenuIdByRoleId(roleId int32) []int32
	GetMenusByPid(pid int32) []*models.SysMenu
	GetMenuIdByUserId(userId int32) []int32
	GetMenuPermByUserId(userId int32) []string
	SaveMenuApi(menuId int32, apiId int32) bool
	DelMenuApi(menuId int32) bool
	GetApiIdByMenuId(menuId int32) []int32
}

type SysMenuRepo struct {
	Log      logger.ILogger `inject:""`
	BaseRepo BaseRepo       `inject:"inline"`
}

func (repo *SysMenuRepo) GetMenuById(id int32) *models.SysMenu {
	var menu models.SysMenu
	if err := repo.BaseRepo.FirstByID(&menu, int(id)); err != nil {
		repo.Log.Errorf("获取菜单数据失败", err)
	}
	return &menu
}

func (repo *SysMenuRepo) InsertMenu(menu *models.SysMenu) bool {
	if err := repo.BaseRepo.Create(menu); err != nil {
		repo.Log.Errorf("新增菜单失败", err)
		return false
	}
	return true
}

func (repo *SysMenuRepo) UpdateMenu(menu *models.SysMenu) bool {
	if err := repo.BaseRepo.Source.DB().Model(&menu).Update(menu).Error; err != nil {
		repo.Log.Errorf("更新菜单失败", err)
		return false
	}
	return true
}

func (repo *SysMenuRepo) DeleteMenu(id int32) bool {
	menu := models.SysMenu{}
	where := &models.SysMenu{MenuId: id}
	if count, err := repo.BaseRepo.DeleteByWhere(&menu, where); err != nil {
		repo.Log.Errorf("删除菜单失败", err)
		return false
	} else {
		return count > 0
	}
}

func (repo *SysMenuRepo) ListMenu(page, size int32, total *int32, where interface{}) []*models.SysMenu {
	var menus []*models.SysMenu
	if err := repo.BaseRepo.GetPages(&models.SysMenu{}, &menus, page, size, total, where); err != nil {
		repo.Log.Errorf("获取菜单列表失败", err)
	}
	return menus
}

func (repo *SysMenuRepo) GetMenuByRoleId(roleId int32) []*models.SysMenu {
	var menus []*models.SysMenu
	err := repo.BaseRepo.Source.DB().Raw("select m.* from t_sys_menu m, t_sys_role_menu rm where m.menu_id = rm.menu_id and rm.role_id = ?", roleId).Find(&menus).Error
	if err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return menus
}

func (repo *SysMenuRepo) GetMenuIdByRoleId(roleId int32) []int32 {
	var menuIds []int32
	where := map[string]int32{"role_id": roleId}
	if err := repo.BaseRepo.Source.DB().Table("t_sys_role_menu").Find(&where, &menuIds, "menu_id"); err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return menuIds
}

func (repo *SysMenuRepo) GetMenuIdByUserId(userId int32) []int32 {
	var menuIds []int32
	rows, err := repo.BaseRepo.Source.DB().Raw("select distinct rm.menu_id from t_sys_user_role ur, t_sys_role_menu rm where ur.role_id = rm.role_id and m.is_del = 0 and ur.user_id = ?", userId).Rows()
	if err != nil {
		repo.Log.Errorf("查询数据失败", err)
		return menuIds
	}
	defer rows.Close()
	for rows.Next() {
		var menuId int32
		_ = rows.Scan(&menuId)
		menuIds = append(menuIds, menuId)
	}
	return menuIds
}

func (repo *SysMenuRepo) GetMenuPermByUserId(userId int32) []string {
	var perms []string
	rows, err := repo.BaseRepo.Source.DB().Raw("select distinct m.perm from t_sys_menu m, t_sys_user_role ur, t_sys_role_menu rm where ur.role_id = rm.role_id and rm.menu_id = m.menu_id and m.is_del = 0 and ur.user_id = ?", userId).Rows()
	if err != nil {
		repo.Log.Errorf("查询数据失败", err)
		return perms
	}
	defer rows.Close()
	for rows.Next() {
		var perm string
		_ = rows.Scan(&perm)
		perms = append(perms, perm)
	}
	return perms
}

func (repo *SysMenuRepo) GetMenusByPid(pid int32) []*models.SysMenu {
	var menus []*models.SysMenu
	where := &models.SysMenu{ParentMenuId: pid}
	if err := repo.BaseRepo.Source.DB().Where(where).Find(&menus); err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return menus
}

func (repo *SysMenuRepo) SaveMenuApi(menuId int32, apiId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("insert into t_sys_menu_api(menu_id, app_id) values (?, ?)", menuId, apiId).Error; err != nil {
		repo.Log.Errorf("写入数据失败", err)
		return false
	}
	return true
}

func (repo *SysMenuRepo) DelMenuApi(menuId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("delete from t_sys_menu_api where menu_id = ?", menuId).Error; err != nil {
		repo.Log.Errorf("删除数据失败", err)
		return false
	}
	return true
}

func (repo *SysMenuRepo) GetApiIdByMenuId(menuId int32) []int32 {
	var apiIds []int32
	where := map[string]int32{"menu_id": menuId}
	if err := repo.BaseRepo.Find(&where, &apiIds, "menu_id"); err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return apiIds
}
