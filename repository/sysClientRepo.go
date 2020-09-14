package repository

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
)

type ISysClientRepo interface {
	GetClientById(id int32) *models.SysClient
	InsertClient(client *models.SysClient) bool
	UpdateClient(client *models.SysClient) bool
	DeleteClient(id int32) bool
	ListClient(page, size int32, total *int32, where interface{}) []*models.SysClient
	GetClientIp(id int32) []string
	SaveClientIpApi(clientId int32, ip string, appId int32) bool
	DelClientIp(clientId int32, ip string) bool
	GetClientIpApi(clientId int32, ip string) []int32
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

func (repo *SysClientRepo) GetClientIp(id int32) []string {
	var ips []string
	rows, err := repo.BaseRepo.Source.DB().Raw("select distinct ip from t_sys_client_api where client_id = ?", id).Rows()
	if err != nil {
		repo.Log.Errorf("查询数据失败", err)
		return ips
	}
	defer rows.Close()
	for rows.Next() {
		var ip string
		_ = rows.Scan(&ip)
		ips = append(ips, ip)
	}
	return ips
}

func (repo *SysClientRepo) SaveClientIpApi(clientId int32, ip string, appId int32) bool {
	if err := repo.BaseRepo.Source.DB().Exec("insert into t_sys_client_api(client_id, ip, app_id) values (?, ?, ?)", clientId, ip).Error; err != nil {
		repo.Log.Errorf("写入数据失败", err)
		return false
	}
	return true
}

func (repo *SysClientRepo) DelClientIp(clientId int32, ip string) bool {
	if err := repo.BaseRepo.Source.DB().Exec("delete from t_sys_client_api where client_id = ? and ip = ?", clientId, ip).Error; err != nil {
		repo.Log.Errorf("删除数据失败", err)
		return false
	}
	return true
}

func (repo *SysClientRepo) GetClientIpApi(clientId int32, ip string) []int32 {
	var appIds []int32
	where := vo.SysClientApiVo{
		ClientId: clientId,
		Ip:       ip,
	}
	if err := repo.BaseRepo.PluckList(&vo.SysClientApiVo{}, &where, &appIds, "app_id").Error; err != nil {
		repo.Log.Errorf("查询数据失败", err)
	}
	return appIds
}
