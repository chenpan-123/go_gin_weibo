package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go_gin_weibo/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))
	r.Use(sessions.Sessions("mysession", store))

	{
		//注册：
		r.GET("/register", controllers.RegisterGet)
		r.POST("/register", controllers.RegisterPost)

		//登录：
		r.GET("/login", controllers.LoginGet)
		r.POST("/login", controllers.LoginPost)

		//首页
		r.GET("/", controllers.HomeGet)

		//退出
		r.GET("/exit", controllers.ExitGet)
	}

	return r
}
