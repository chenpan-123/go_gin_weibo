package main

import (
	"go_gin_weibo/databases"
	"go_gin_weibo/routers"
)

func main() {
	databases.InitMysql()
	r := routers.InitRouter()
	//静态资源
	r.Static("/static", "./static")
	r.Run(":8081")
}
