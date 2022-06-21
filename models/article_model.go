package models

import (
	"fmt"
	"go_gin_weibo/config"
	"go_gin_weibo/databases"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

//----数据处理----
//----添加文章----
func AddArticle(article Article) (int64, error) {
	i, err := InsertArticle(article)
	SetArticleRowsNum()
	return i, err

}

//----数据库操作----

//插入一篇文章
func InsertArticle(article Article) (int64, error) {
	return databases.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

//----查询文章----

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	fmt.Println("-------->page", page)
	//从配置文件中获取每页的文章数量
	return QueryArticleWithPage(page, config.NUM)

}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m,n
	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据
注意limit前面有where

*/

func QueryArticleWithPage(page int, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article" + sql
	rows, err := databases.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

//----翻页----

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum

}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := databases.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num

}

//设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()

}
