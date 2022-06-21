package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_weibo/models"
	"net/http"
	"strconv"
)

func HomeGet(c *gin.Context) {

	//获取session，判断用户是否登录
	islogin := GetSession(c)

	tag := c.Query("tag")
	fmt.Println("tag:", tag)
	page, _ := strconv.Atoi(c.Query("page"))
	var artList []models.Article

	if page <= 0 {
		page = 1
	}

	artList, _ = models.FindArticleWithPage(page)

	homeFooterPageCode := models.ConfigHomeFooterPageCode(page)
	html := models.MakeHomeBlock(artList, islogin)

	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin, "Content": html, "HasFooter": true, "PageCode": homeFooterPageCode})

}
