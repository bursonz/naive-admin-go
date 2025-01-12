package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"naive-admin-go/api"
	"naive-admin-go/middleware"
)

func Init(r *gin.Engine) {
	// 使用 cookie 存储会话数据
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	r.Use(middleware.Cors())

	r.POST("/auth/login", api.Auth.Login)
	r.GET("/auth/captcha", api.Auth.Captcha) // 获取验证码

	r.Use(middleware.Jwt())
	r.POST("/auth/logout", api.Auth.Logout)
	r.POST("/auth/password", api.Auth.Logout)             //TODO
	r.GET("/auth/refresh/:expire", api.Auth.RefreshToken) // 刷新token有效期

	r.GET("/user", api.User.List)
	r.POST("/user", api.User.Add)
	r.DELETE("/user/:id", api.User.Delete)
	r.PATCH("/user/password/reset/:id", api.User.Update)
	r.PATCH("/user/:id", api.User.Update)
	r.PATCH("/user/profile/:id", api.User.Profile)
	r.GET("/user/detail", api.User.Detail)

	r.GET("/role", api.Role.List)
	r.POST("/role", api.Role.Add)
	r.PATCH("/role/:id", api.Role.Update)
	r.DELETE("/role/:id", api.Role.Delete)
	r.PATCH("/role/users/add/:id", api.Role.AddUser)
	r.PATCH("/role/users/remove/:id", api.Role.RemoveUser)
	r.GET("/role/page", api.Role.ListPage)
	r.GET("/role/permissions/tree", api.Role.PermissionsTree)

	r.POST("/permission", api.Permissions.Add)
	r.PATCH("/permission/:id", api.Permissions.PatchPermission)
	r.DELETE("/permission/:id", api.Permissions.Delete)
	r.GET("/permission/tree", api.Permissions.List)
	r.GET("/permission/menu/tree", api.Permissions.List)

	r.GET("/station", api.Station.List)
	r.POST("/station", api.Station.Add)
	r.PATCH("/station/:id", api.Station.Update)
	r.DELETE("/station/:id", api.Station.Delete)

	r.GET("/lock", api.Lock.List)
	r.POST("/lock", api.Lock.Add)
	r.PATCH("/lock/:id", api.Lock.Update)
	r.DELETE("/lock/:id", api.Lock.Delete)
}
