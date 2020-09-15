package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
)

type ISysRoleRepo interface {
	GetRoleById(id int32) *models.SysRole
	InsertRole(role *models.SysRole) bool
	UpdateRole(role *models.SysRole) bool
	DeleteRole(id int32) bool
	ListRole(page, size int32, total *int32, where interface{}) []*models.SysRole
	AllRole() []*models.SysRole
	GetRoleByUserId(userId int32) []*models.SysRole
	GetRoleIdByUserId(userId int32) []int32
	CountUserByRoleId(roleId int32) int32
	SaveRoleMenu(roleId, menuId int32) bool
	DelRoleMenu(roleId int32) bool
}

type SysRoleRepo struct {
	Log      logger.ILogger `inject:""`
	BaseRepo BaseRepo       `inject:"inline"`
}

func (repo *SysRoleRepo) GetRoleById(id int32) *models.SysRole {
	var role models.SysRole
	if err := repo.BaseRepo.FirstByID(&role, int(id)); err != nil {
		repo.Log.Errorf("获取角色数据失败", err)
	}
	return &role
}

func (repo *SysRoleRepo) InsertRole(role *models.SysRole) bool {
	if err := repo.BaseRepo.Create(role); err != nil {
		repo.Log.Errorf("新增角色失败", err)
		return false
	}
	return true
}

func (repo *SysRoleRepo) UpdateRole(role *models.SysRole) bool {
	if err := repo.BaseRepo.Source.DB().Model(&role).Update(role).Error; err != nil {
		repo.Log.Errorf("更新角色失败", err)
		return false
	}
	return true
}

func (repo *SysRoleRepo) DeleteRole(id int32) bool {
	role := models.SysRole{}
	where := &models.SysRole{RoleId: id}
	if count, err := repo.BaseRepo.DeleteByWhere(&role, where); err != nil {
		repo.Log.Errorf("删除角色失败", err)
		return false
	} else {
		return count > 0
	}
}

func (repo *SysRoleRepo) ListRole(page, size int32, total *int32, where interface{}) []*models.SysRole {
	var roles []*models.SysRole
	if err := repo.BaseRepo.GetPages(&models.SysRole{}, &roles, page, size, total, where); err != nil {
		repo.Log.Errorf("获取角色列表失败", err)
	}
	return roles
}

func (repo *SysRoleRepo) AllRole() []*models.SysRole {
	var roles []*models.SysRole
	if err := repo.BaseRepo.Source.DB().Find(&roles); err != nil {
		repo.Log.Errorf("获取角色失败", err)
	}
	return roles
}

func (repo *SysRoleRepo) GetRoleByUserId(userId int32) []*models.SysRole {
	var roles []*models.SysRole
	err := repo.BaseRepo.Source.DB().Raw("select r.* from t_sys_role r, t_sys_user_role ur where r.role_id = ur.role_id and r.is_del = 0 and ur.user_id = ?", userId).Find(&roles).Error
	if err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return roles
}

func (repo *SysRoleRepo) GetRoleIdByUserId(userId int32) []int32 {
	var roleIds []int32
	if err := repo.BaseRepo.Source.DB().Table("t_sys_user_role").Where("user_id = ?", userId).Pluck("role_id", &roleIds).Error; err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return roleIds
}

// 统计该角色关联的用户数
func (repo *SysRoleRepo) CountUserByRoleId(roleId int32) int32 {
	var count int32
	if err := repo.BaseRepo.Source.DB().Table("t_sys_user_role").Where("role_id", roleId).Count(&count).Error; err != nil {
		repo.Log.Errorf("查询计数失败", err)
	}
	return count
}

func (repo *SysRoleRepo) SaveRoleMenu(roleId, menuId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("insert into t_sys_role_menu(role_id, menu_id) values (?, ?)", roleId, menuId).Error; err != nil {
		repo.Log.Errorf("写入数据失败", err)
		return false
	}
	return true
}

func (repo *SysRoleRepo) DelRoleMenu(roleId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("delete from t_sys_role_menu where role_id = ?", roleId).Error; err != nil {
		repo.Log.Errorf("删除数据失败", err)
		return false
	}
	return true
}
