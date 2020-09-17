package routers

import (
	"github.com/PGshen/june/common/datasource"
	"github.com/PGshen/june/common/logger"
	"github.com/PGshen/june/common/middleware"
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
	var menu web.SysMenuWeb
	var role web.SysRoleWeb
	var user web.SysUserWeb
	var myjwt middleware.Jwt
	db := datasource.Db{}
	zap := logger.Logger{}
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &api},
		&inject.Object{Value: &client},
		&inject.Object{Value: &menu},
		&inject.Object{Value: &role},
		&inject.Object{Value: &user},
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &myjwt},
		&inject.Object{Value: &repository.BaseRepo{}},
		&inject.Object{Value: &repository.SysApiRepo{}},
		&inject.Object{Value: &service.SysApiService{}},
		&inject.Object{Value: &repository.SysClientRepo{}},
		&inject.Object{Value: &service.SysClientService{}},
		&inject.Object{Value: &repository.SysMenuRepo{}},
		&inject.Object{Value: &service.SysMenuService{}},
		&inject.Object{Value: &repository.SysRoleRepo{}},
		&inject.Object{Value: &service.SysRoleService{}},
		&inject.Object{Value: &repository.SysUserRepo{}},
		&inject.Object{Value: &service.SysUserService{}},
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
	var authMiddleware = myjwt.GinJWTMiddlewareInit(middleware.Authorizator)
	r.NoRoute(authMiddleware.MiddlewareFunc(), middleware.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.Use(authMiddleware.MiddlewareFunc())
	sysApi := r.Group("")
	{
		sysApi.GET("/api", api.GetApiById)
		sysApi.POST("/apis/list", api.ListApi)
		sysApi.POST("/api", api.SaveApi)
		sysApi.PUT("/api/:id", api.EditApi)
		sysApi.DELETE("/api/:id", api.DelApi)
		sysApi.GET("/api/tree", api.ApiTree)
		sysApi.GET("/api/tree/:id", api.ApiTreeById)
	}
	sysClient := r.Group("")
	{
		sysClient.GET("/client", client.GetClientById)
		sysClient.POST("/clients/list", client.ListClient)
		sysClient.POST("/client", client.SaveClient)
		sysClient.PUT("/client/:id", client.EditClient)
		sysClient.DELETE("/client/:id", client.DelClient)

		sysClient.GET("/clients/ip/:id", client.GetClientIp)
		sysClient.POST("/clients/ip", client.SaveClientIp)
		sysClient.DELETE("/clients/ip/:id", client.DelClientIp)

		sysClient.POST("/clients/ips/apis/list", client.GetClientIpApi)
		sysClient.POST("/clients/ips/api", client.AuthClientIpApi)
	}
	sysMenu := r.Group("")
	{
		sysMenu.GET("/menu", menu.GetMenuById)
		sysMenu.POST("/menus/list", menu.ListMenu)
		sysMenu.POST("/menu", menu.SaveMenu)
		sysMenu.PUT("/menu/:id", menu.EditMenu)
		sysMenu.DELETE("/menu/:id", menu.DelMenu)

		sysMenu.GET("/menus/tree/:id", menu.GetMenuTreeById)
		sysMenu.GET("/menus/trees/all", menu.GetMenuTree)

		sysMenu.GET("/menus/api/:id", menu.GetMenuApiById)
		sysMenu.PUT("/menus/api", menu.AuthMenuApi)
	}
	sysRole := r.Group("")
	{
		sysRole.GET("/role", role.GetRoleById)
		sysRole.POST("/roles/list", role.ListRole)
		sysRole.GET("/roles/all", role.GetAllRole)
		sysRole.POST("/role", role.SaveRole)
		sysRole.PUT("/role/:id", role.EditRole)
		sysRole.DELETE("/role/:id", role.DelRole)

		sysRole.GET("/roles/menu/:id", role.GetRoleMenu)
		sysRole.PUT("/roles/auth", role.AuthRoleMenu)
	}
	sysUser := r.Group("")
	{
		sysUser.GET("/user", user.GetUserById)
		sysUser.POST("/users/list", user.ListUser)
		sysUser.POST("/user", user.SaveUser)
		sysUser.PUT("/user/:id", user.EditUser)
		sysUser.DELETE("/user/:id", user.DelUser)

		sysUser.GET("/users/role/:id", user.GetUserRole)
		sysUser.PUT("/users/ban/:id", user.EnableUser)
		sysUser.GET("/users/login/:loginName", user.GetUserByLoginName)
		sysUser.POST("/users/now", user.NowUser)
	}
}
