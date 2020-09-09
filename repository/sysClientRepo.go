package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
)

type ISysClientRepo interface {
	GetClientById(id int32) *models.SysClient
	InsertClient(client *models.SysClient) bool
	UpdateClient(client *models.SysClient) bool
	DeleteClient(id int32) bool
	ListClient(page, size int32, total *int32, where interface{}) []*models.SysClient
}

type SysClientRepo struct {
	Log      logger.ILogger `inject:""`
	BaseRepo BaseRepo       `inject:"inline"`
}

func (repo *SysClientRepo) GetClientById(id int32) *models.SysClient {
	var client models.SysClient
	if err := repo.BaseRepo.FirstByID(&client, int(id)); err != nil {
		repo.Log.Errorf("找不到记录", err)
		return nil
	}
	return &client
}

func (repo *SysClientRepo) InsertClient(client *models.SysClient) bool {
	if err := repo.BaseRepo.Create(client); err != nil {
		repo.Log.Errorf("插入数据失败", err)
		return false
	}
	return true
}

func (repo *SysClientRepo) UpdateClient(client *models.SysClient) bool {
	if err := repo.BaseRepo.Source.DB().Model(&client).Update(client).Error; err != nil {
		repo.Log.Errorf("更新数据失败", err)
		return false
	}
	return true
}

func (repo *SysClientRepo) DeleteClient(id int32) bool {
	if count, err := repo.BaseRepo.DeleteByWhere(&models.SysClient{}, &models.SysClient{ClientId: id}); err != nil {
		repo.Log.Errorf("删除数据失败", err)
		return false
	} else {
		return count > 0
	}
}

func (repo *SysClientRepo) ListClient(page, size int32, total *int32, where interface{}) []*models.SysClient {
	var clients []*models.SysClient
	if err := repo.BaseRepo.GetPages(&models.SysClient{}, &clients, page, size, total, where); err != nil {
		repo.Log.Errorf("获取数据失败", err)
	}
	return clients
}
