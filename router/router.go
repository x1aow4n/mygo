package router

import (
	"github.com/gin-gonic/gin"
	"mygo/handler"
)

func InitRouter(r *gin.Engine) {
	r.POST("/user", handler.UserRegister)
}
