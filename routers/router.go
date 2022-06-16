package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_weibo/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	//注册
	r.GET("/register", controllers.RegisterGet)
	return r
}
