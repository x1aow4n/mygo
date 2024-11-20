package main

import (
	"github.com/gin-gonic/gin"
	"mygo/db"
	"mygo/router"
)

func main() {
	r := gin.Default()
	db.InitDb()
	router.InitRouter(r)
	err := r.Run("localhost:8080")
	if err != nil {
		panic("服务启动失败")
	}
}
