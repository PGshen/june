package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
)

type ISysApiRepo interface {
	GetApiById(id int) *models.SysApi
	InsertApi(api *models.SysApi) bool
	UpdateApi(api *models.SysApi) bool
	DeleteApi(id int) bool
}

// 依赖注入
type SysApiRepo struct {
	Log      logger.ILogger `inject:""`
	BaseRepo BaseRepo       `inject:"inline"`
}

// 通过ID查找
func (apiRepo *SysApiRepo) GetApiById(id int) *models.SysApi {
	var api models.SysApi
	if err := apiRepo.BaseRepo.FirstByID(&api, id); err != nil {
		apiRepo.Log.Errorf("找不到记录", err)
		return nil
	}
	return &api
}

// 保存
func (apiRepo *SysApiRepo) InsertApi(api *models.SysApi) bool {
	if err := apiRepo.BaseRepo.Create(api); err != nil {
		apiRepo.Log.Errorf("新建AP接口失败", err)
		return false
	}
	return true
}

// 更新
func (apiRepo *SysApiRepo) UpdateApi(api *models.SysApi) bool {
	//使用事务同时更新用户数据和角色数据
	if err := apiRepo.BaseRepo.Save(api).Error; err != nil {
		apiRepo.Log.Errorf("更新API接口失败", err)
		return false
	}
	return true
}

// 删除
func (apiRepo *SysApiRepo) DeleteApi(apiId int) bool {
	api := models.SysApi{}
	if err := apiRepo.BaseRepo.DeleteByID(api, apiId); err != nil {
		apiRepo.Log.Errorf("删除API接口失败", err)
		return false
	}
	return true
}
