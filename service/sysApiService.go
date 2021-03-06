package service

import (
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/req"
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/common/utils"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysApiService interface {
	GetApi(c *gin.Context, id int)
	SaveApi(c *gin.Context, sysApi *models.SysApi)
	UpdateApi(c *gin.Context, sysApi *models.SysApi)
	DeleteApi(c *gin.Context, id int)
	ListApi(c *gin.Context, reqCond *req.ReqCond)
	GetApiTrees(c *gin.Context)
	GetApiTreeById(c *gin.Context, id int)
	GetApiTree(id int) *models.SysApiTree
	CutApiTree(flag bool, trees []models.SysApiTree, menuApiIds []int32) []models.SysApiTree
}

type SysApiService struct {
	Repo repository.ISysApiRepo `inject:""`
	Log  logger.ILogger         `inject:""`
}

// 获取API
func (apiService *SysApiService) GetApi(c *gin.Context, id int) {
	apiService.Log.Infof("Get api: id = %s", id)
	data := apiService.Repo.GetApiById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Api, data)
	}
}

// 保存API
func (apiService *SysApiService) SaveApi(c *gin.Context, sysApi *models.SysApi) {
	apiService.Log.Info("Save api: api = %s", sysApi)
	if apiService.Repo.InsertApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0313, nil)
	}
}

// 更新API
func (apiService *SysApiService) UpdateApi(c *gin.Context, sysApi *models.SysApi) {
	apiService.Log.Info("Update api: api = %s", sysApi)
	if apiService.Repo.UpdateApi(sysApi) {
		resp.RespB200(c, bcode.Api, sysApi)
	} else {
		resp.RespB406(c, bcode.Api, ecode.P0312, nil)
	}
}

// 删除Api
func (apiService *SysApiService) DeleteApi(c *gin.Context, id int) {
	apiService.Log.Info("Delete api: id = %s", id)
	if apiService.Repo.DeleteApi(id) {
		resp.RespB200(c, bcode.Api, "")
	} else {
		resp.RespB406s(c, bcode.Api, ecode.P0302, "", nil)
	}
}

// 获取Api列表
func (apiService *SysApiService) ListApi(c *gin.Context, reqCond *req.ReqCond) {
	apiService.Log.Info("List api, condition = %s", reqCond)
	page := reqCond.Page
	size := reqCond.Size
	var total int32
	where := utils.GetFilter(reqCond.Filter)
	apis := apiService.Repo.ListApi(page, size, &total, where)
	res := make(map[string]interface{})
	res["records"] = apis
	res["total"] = total
	resp.RespB200(c, bcode.Api, res)
}

// 完整API树
func (apiService *SysApiService) GetApiTrees(c *gin.Context) {
	apiService.Log.Info("Get api tree")
	apiTree := apiService.GetApiTree(1)
	if apiTree == nil {
		resp.RespB406(c, bcode.Api, ecode.P0301, nil)
	} else {
		res := []models.SysApiTree{*apiTree}
		resp.RespB200(c, bcode.Api, res)
	}
}

// 指定ID的API树
func (apiService *SysApiService) GetApiTreeById(c *gin.Context, id int) {
	apiService.Log.Info("Get api tree by id, id = %s", id)
	// todo 检查ID
	apiTree := apiService.GetApiTree(id)
	if apiTree == nil {
		resp.RespB406(c, bcode.Api, ecode.P0301, nil)
	} else {
		resp.RespB200(c, bcode.Api, apiTree)
	}
}

func (apiService *SysApiService) GetApiTree(id int) *models.SysApiTree {
	var apiTree = models.SysApiTree{}
	api := apiService.Repo.GetApiById(id)
	if api == nil {
		return nil
	}
	apiTree.SysApi = *api
	sysApis := apiService.Repo.GetApiByPid(id)
	// 递归查询子节点
	if sysApis == nil || len(sysApis) == 0 {
		apiTree.Children = []models.SysApiTree{}
	} else {
		for child := range sysApis {
			apiTree.Children = append(apiTree.Children, *apiService.GetApiTree(int(sysApis[child].ApiId)))
		}
	}
	return &apiTree
}

// 裁剪
func (apiService *SysApiService) CutApiTree(flag bool, trees []models.SysApiTree, menuApiIds []int32) []models.SysApiTree {
	for e := range trees {
		if trees[e].Children != nil {
			trees[e].Children = apiService.CutApiTree(flag, trees[e].Children, menuApiIds)
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
