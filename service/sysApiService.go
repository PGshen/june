package service

import (
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/repository"
	"github.com/gin-gonic/gin"
)

type ISysApiService interface {
	GetApi(c *gin.Context, id int)
}

type SysApiService struct {
	Repo repository.ISysApiRepo `inject:""`
}

func (apiService *SysApiService) GetApi(c *gin.Context, id int) {
	data := apiService.Repo.GetApiById(id)
	if data == nil {
		resp.RespB406s(c, bcode.Api, ecode.P0301, "", nil)
	} else {
		resp.RespB200(c, bcode.Api, data)
	}
}
