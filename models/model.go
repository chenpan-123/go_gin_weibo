package models

import (
	"fmt"
	"go_gin_weibo/databases"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Status    int // 0 正常，1 删除
	Creattime int64
}

//------------------数据库操作-----------------------

//插入
func InsertUser(user User) (int64, error) {
	return databases.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Creattime)
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := databases.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}
