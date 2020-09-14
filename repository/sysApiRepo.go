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
	ListApi(page int32, size int32, total *int32, where interface{}) []*models.SysApi
	GetApiByPid(pid int) []*models.SysApi
	GetApiByUserId(userId int32) []*models.SysApi
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

func (apiRepo *SysApiRepo) GetApiByPid(pid int) []*models.SysApi {
	var apis []*models.SysApi
	where := &models.SysApi{ParentApiId: int32(pid)}
	if err := apiRepo.BaseRepo.Source.DB().Where(where).Find(&apis); err != nil {
		apiRepo.Log.Error("获取API失败", err)
	}
	return apis
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
	//if err := apiRepo.BaseRepo.Source.DB().Model(&models.SysApi{}).Where("api_id = ?", api.ApiId).Update(api).Error; err != nil {
	// api 如果带了主键则不需要再查询一次
	if err := apiRepo.BaseRepo.Source.DB().Model(&api).Update(api).Error; err != nil {
		apiRepo.Log.Errorf("更新API接口失败", err)
		return false
	}
	return true
}

// 删除
func (apiRepo *SysApiRepo) DeleteApi(apiId int) bool {
	api := models.SysApi{}
	where := &models.SysApi{ApiId: int32(apiId)}
	if count, err := apiRepo.BaseRepo.DeleteByWhere(&api, where); err != nil {
		apiRepo.Log.Errorf("删除API接口失败", err)
		return false
	} else {
		return count > 0
	}
}

func (apiRepo *SysApiRepo) ListApi(page int32, size int32, total *int32, where interface{}) []*models.SysApi {
	var apis []*models.SysApi
	if err := apiRepo.BaseRepo.GetPages(&models.SysApi{}, &apis, page, size, total, where); err != nil {
		apiRepo.Log.Error("获取API列表失败：", err)
	}
	return apis
}

func (apiRepo *SysApiRepo) GetApiByUserId(userId int32) []*models.SysApi {
	var apis []*models.SysApi
	err := apiRepo.BaseRepo.Source.DB().Raw("SELECT distinct a.* from t_sys_user_role ur, t_sys_role_menu rm, t_sys_menu_api ma, t_sys_api a where a.is_del = 0 AND ur.role_id = rm.role_id AND rm.menu_id = ma.menu_id AND ur.user_id = ?", userId).Find(&apis).Error
	if err != nil {
		apiRepo.Log.Errorf("查询数据失败", err)
	}
	return apis
}
