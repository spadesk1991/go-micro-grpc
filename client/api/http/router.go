package httpapi

import (
	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/go-micro-grpc/client/api/http/controllers"
)

func InitRouter(router *gin.Engine) {
	router.GET("/order", controllers.GetOrderById)
}
