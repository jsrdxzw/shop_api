package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/user-web/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("user")
	userGroup.GET("list", api.GetUserList)
}
