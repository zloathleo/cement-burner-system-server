package context

import (
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/my"
)

func Init(router *gin.Engine) {
	router.Use(newCrosMiddleware())

	pubGroup := router.Group("/")
	myGroup := router.Group("/my", newAuthMiddleware())

	my.InitDashboard(pubGroup, myGroup)
}
