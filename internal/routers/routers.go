package routers

import (
	"github.com/CLannadZSY/TaskSupervisor/internal/routers/user"
	"github.com/gin-gonic/gin"
)

func init() {
}

// 注册路由
func RegisterRouters() {

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//router.StaticFS()

	routerGroup := router.Group("/api")
	userGroup := routerGroup.Group("/user")
	{
		userGroup.POST("/login", user.LoginUser)
		userGroup.POST("/registry", user.RegistryUser)
	}

	router.Run()
}
