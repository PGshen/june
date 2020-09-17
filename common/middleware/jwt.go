package middleware

import (
	"github.com/PGshen/june/common/resp"
	"github.com/PGshen/june/common/returncode/bcode"
	"github.com/PGshen/june/common/returncode/ecode"
	"github.com/PGshen/june/common/returncode/hcode"
	"github.com/PGshen/june/common/returncode/tcode"
	"github.com/PGshen/june/common/setting"
	"github.com/PGshen/june/common/utils"
	"github.com/PGshen/june/models"
	"github.com/PGshen/june/models/vo"
	"github.com/PGshen/june/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type Jwt struct {
	Log         log.Logger
	UserService service.ISysUserService `inject:""`
	RoleService service.ISysRoleService `inject:""`
}

// JwtAuthorizator 定义身份授权事件类型
type JwtAuthorizator func(data interface{}, c *gin.Context) bool

//app 程序配置
var app = setting.Config.App

//GinJWTMiddlewareInit 初始化中间件
func (j *Jwt) GinJWTMiddlewareInit(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: app.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*vo.UserPermVo); ok {
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userId":    v.UserId,
					"loginName": v.LoginName,
					"perms":     utils.Strlist2str(v.Perms),
					"roles":     utils.Intlist2str(v.Roles),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			userInfoVo := jwt.ExtractClaims(c)
			userId := int32(userInfoVo["userId"].(float64))
			loginName := userInfoVo["loginName"].(string)
			perms := strings.Split(userInfoVo["perms"].(string), ",")
			roles := utils.Intlist2strlist(strings.Split(userInfoVo["roles"].(string), ","))
			//Set the identity
			return &vo.UserPermVo{
				UserId:    userId,
				LoginName: loginName,
				Roles:     roles,
				Perms:     perms,
			}
		},
		// 登录逻辑
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var user models.SysUser
			if err := c.ShouldBind(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			loginName := user.LoginName
			password := user.Password
			if j.UserService.CheckUser(loginName, password) {
				userInfo := j.UserService.GetUserInfoByLoginName(loginName)
				userId := userInfo.User.UserId
				roles := userInfo.Roles
				perms := userInfo.Permissions
				return &vo.UserPermVo{
					UserId:    userId,
					LoginName: loginName,
					Roles:     roles,
					Perms:     perms,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			resp.RespBusiData(c, hcode.Ok, tcode.Business, bcode.Auth, ecode.P0508, "授权失败", nil)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			data := map[string]interface{}{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			}
			resp.RespB200(c, bcode.Auth, data)
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			data := map[string]interface{}{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			}
			resp.RespB200m(c, bcode.Auth, "刷新成功", data)
		},
		LogoutResponse: func(c *gin.Context, code int) {
			resp.RespB200m(c, bcode.Auth, "已退出登录", nil)
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		j.Log.Fatal("JWT Error:" + err.Error())
	}
	return
}

// 权限检查 todo 优化逻辑
func Authorizator(data interface{}, c *gin.Context) bool {
	uri := c.Request.RequestURI
	method := c.Request.Method
	if v, ok := data.(*vo.UserPermVo); ok {
		for _, perm := range v.Perms {
			if permMatch(perm, method, uri) {
				return true
			}
			return true
		}
	}
	return false
}

//NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	resp.RespBusiData(c, hcode.NotFound, tcode.Business, bcode.BusinessAbnormal, ecode.P0602, "404", nil)
}

// 权限与请求匹配
func permMatch(perm, method, uri string) bool {
	// 去除?后的参数
	if index := strings.Index(uri, "?"); index > 0 {
		uri = uri[:index]
	}
	// perm ->  method#uri
	if !strings.Contains(perm, ":") {
		return false
	}
	permArr := strings.Split(perm, ":")
	if permArr[0] != method {
		return false
	}
	sUris := strings.Split(permArr[1], "/")
	dUris := strings.Split(uri, "/")
	if len(sUris) != len(dUris) {
		return false
	}
	for e := range sUris {
		if sUris[e] != dUris[e] && strings.Index(sUris[e], "{") != 0 && strings.LastIndex(sUris[e], "}") != len(sUris)-1 {
			return false
		}
	}
	return true
}
