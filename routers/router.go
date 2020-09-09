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
	var client web.SysClientWeb
	db := datasource.Db{}
	zap := logger.Logger{}
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &api},
		&inject.Object{Value: &client},
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &repository.BaseRepo{}},
		&inject.Object{Value: &repository.SysApiRepo{}},
		&inject.Object{Value: &service.SysApiService{}},
		&inject.Object{Value: &repository.SysClientRepo{}},
		&inject.Object{Value: &service.SysClientService{}},
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
		sysApi.GET("/api", api.GetApiById)
		sysApi.POST("/apis", api.ListApi)
		sysApi.POST("/api", api.SaveApi)
		sysApi.PUT("/api/:id", api.EditApi)
		sysApi.DELETE("/api/:id", api.DelApi)
		sysApi.GET("/api/tree", api.ApiTree)
		sysApi.GET("/api/tree/:id", api.ApiTreeById)
	}
	sysClient := r.Group("")
	{
		sysClient.GET("/client", client.GetClientById)
		sysClient.POST("/clients", client.ListClient)
		sysClient.POST("/client", client.SaveClient)
		sysClient.PUT("/client/:id", client.EditClient)
		sysClient.DELETE("/client/:id", client.DelClient)
	}
}
