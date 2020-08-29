package routers

import (
	"github.com/PGshen/june/common/datasource"
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/repository"
	"github.com/PGshen/june/service"
	"github.com/PGshen/june/web"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//gin.SetMode(setting.Config.App.RunMode)
	Configure(r)
	return r
}

func Configure(r *gin.Engine) {
	//inject declare
	var api web.SysApiWeb
	db := datasource.Db{}
	zap := logger.Logger{}
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &api},
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &repository.SysApiRepo{}},
		&inject.Object{Value: &service.SysApiService{}},
		&inject.Object{Value: &repository.BaseRepo{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	sysApi := r.Group("")
	{
		sysApi.GET("/api/:id", api.GetApiById)
		sysApi.POST("/api", api.SaveApi)
		sysApi.PUT("/api/:id", api.EditApi)
		sysApi.DELETE("/api/:id", api.DelApi)
		sysApi.GET("/api/tree", api.ApiTree)
		sysApi.GET("/api/tree/:id", api.ApiTreeById)
	}
}
