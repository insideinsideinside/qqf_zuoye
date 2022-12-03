package model

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)
var dsn = "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
var db, _= sql.Open("mysql", dsn)
var u User
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ID       int64
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
func InitDB() {
	var dsn = "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	var db, err = sql.Open("mysql", dsn)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

func AddUser(username, password string) {
	sqlStr := "insert into user(id,name,password) values (?,?,?)"
	_, err := db.Exec(sqlStr, FindID(), username, password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

func SelectUser(username string) bool {
	u.Username = username
	sqlStr := "select id, name,password from user where name=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, username).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	return true
}

func SelectPassword(username string) string {
	sqlStr := "select id, name, password from user where name=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	u.Username = username
	err := db.QueryRow(sqlStr, username).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return u.Password
}
func SelectID(username string) int64 {
	sqlStr := "select id, name, password from user where name=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	u.Username = username
	err := db.QueryRow(sqlStr, username).Scan(&u.ID, u.Username, &u.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return 0
	}
	return u.ID
}
func ChangePassword(st User) {
	sqlStr := "update user set password=? where name=?"
	_, err := db.Exec(sqlStr, st.Password, st.Username)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	log.Println("update success")
}
func FindMessage(u User) {
	sqlStr := "select id, name, password from user where name=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, u.Username).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s password:%s\n", u.ID, u.Username, u.Password)
}
func FindID() int {
	sqlStr := "select id, name, password from user where id >=?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return 0
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	i := 1
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return 0
		}
		i++
	}
	return i
}
