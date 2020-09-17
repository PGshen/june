package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
)

type ISysUserRepo interface {
	GetUserById(id int32) *models.SysUser
	GetUserByLoginName(loginName string) *models.SysUser
	CheckUser(loginName, password string) bool
	InsertUser(user *models.SysUser) bool
	UpdateUser(user *models.SysUser) bool
	DeleteUser(id int32) bool
	ListUser(page, size int32, total *int32, where interface{}) []*models.SysUser
	EnableUser(userId int32) bool
	SaveUserRole(userId, roleId int32) bool
	DelUserRole(userId int32) bool
}

type SysUserRepo struct {
	Log      logger.ILogger `inject:""`
	BaseRepo BaseRepo       `inject:"inline"`
}

func (repo *SysUserRepo) GetUserById(id int32) *models.SysUser {
	var user models.SysUser
	if err := repo.BaseRepo.FirstByID(&user, int(id)); err != nil {
		repo.Log.Errorf("获取用户数据失败", err)
	}
	return &user
}

func (repo *SysUserRepo) GetUserByLoginName(loginName string) *models.SysUser {
	var user models.SysUser
	where := models.SysUser{LoginName: loginName}
	if err := repo.BaseRepo.First(&where, &user, "*"); err != nil {
		repo.Log.Errorf("获取用户数据失败", err)
		return nil
	}
	return &user
}

func (repo *SysUserRepo) CheckUser(loginName, password string) bool {
	var count int
	repo.BaseRepo.Source.DB().Table("t_sys_user").Where("login_name = ? and password = ?", loginName, password).Count(&count)
	return count > 0
}

func (repo *SysUserRepo) InsertUser(user *models.SysUser) bool {
	if err := repo.BaseRepo.Create(user); err != nil {
		repo.Log.Errorf("新增用户失败", err)
		return false
	}
	return true
}

func (repo *SysUserRepo) UpdateUser(user *models.SysUser) bool {
	if err := repo.BaseRepo.Source.DB().Model(&user).Update(user).Error; err != nil {
		repo.Log.Errorf("更新用户失败", err)
		return false
	}
	return true
}

func (repo *SysUserRepo) DeleteUser(id int32) bool {
	user := models.SysUser{}
	where := &models.SysUser{UserId: id}
	if count, err := repo.BaseRepo.DeleteByWhere(&user, where); err != nil {
		repo.Log.Errorf("删除用户失败", err)
		return false
	} else {
		return count > 0
	}
}

func (repo *SysUserRepo) ListUser(page, size int32, total *int32, where interface{}) []*models.SysUser {
	var users []*models.SysUser
	if err := repo.BaseRepo.GetPages(&models.SysUser{}, &users, page, size, total, where); err != nil {
		repo.Log.Errorf("获取用户列表失败", err)
	}
	return users
}

func (repo *SysUserRepo) EnableUser(userId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("update t_sys_user set is_enable = abs(is_enable - 1) where user_id = ?", userId).Error; err != nil {
		repo.Log.Errorf("更新数据失败", err)
		return false
	}
	return true
}

func (repo *SysUserRepo) SaveUserRole(userId, roleId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("insert into t_sys_user_role(user_id, role_id) values (?, ?)", userId, roleId).Error; err != nil {
		repo.Log.Errorf("写入数据失败", err)
		return false
	}
	return true
}

func (repo *SysUserRepo) DelUserRole(userId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("delete from t_sys_user_role where user_id = ?", userId).Error; err != nil {
		repo.Log.Errorf("删除数据失败", err)
		return false
	}
	return true
}
