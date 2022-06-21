package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		db, _ = sql.Open("mysql", "root:123456@/go_gin_weibo?charset=utf8")
		CreateTableWithUser()
		CreateTableWithArticle()
	}

}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        createtime INT(10)
        );`
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err1 := result.RowsAffected()
	if err1 != nil {
		log.Println(err)
		return 0, err1
	}
	return count, nil

}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
