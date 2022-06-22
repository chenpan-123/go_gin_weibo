package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_weibo/models"
	"go_gin_weibo/utils"
	"net/http"
	"strconv"
)

func ShowArticleGet(c *gin.Context) {
	//获取session
	isLogin := GetSession(c)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	//渲染html
	c.HTML(http.StatusOK, "show_article.html", gin.H{
		"IsLogin": isLogin,
		"Title":   art.Title,
		"Content": utils.SwitchMarkdownToHtml(art.Content),
	})

}
