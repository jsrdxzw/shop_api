package initialize

import (
	"github.com/gin-gonic/gin"
	"shop_api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
