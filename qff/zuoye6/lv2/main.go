package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	ID     int    `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Gender string `gorm:"column:gender"`
}

func adduser() {
	dsn := "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}) //没有表格则创建一个user表格
	var ID int
	var N, G string
	fmt.Println("请依次输入ID，名字，性别")
	fmt.Scanf("%d %s %s\n", &ID, &N, &G)
	user := User{ID, N, G}
	db.Create(&user)
	fmt.Println("操作完成")
	time.Sleep(time.Second)
} //添加一个用户

func deleteuser() {
	dsn := "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}) //没有表格则创建一个user表格
	var ID int
	fmt.Println("请输入ID")
	fmt.Scanf("%d\n", &ID)
	fmt.Println("操作完成")
	for {
		fmt.Println("返回上一级")
		break
	}
	time.Sleep(time.Second)
	user := User{
		ID: ID,
	}
	db.Delete(&user)

} //删除一个用户
func Return() {
	for {
		x := 0
		fmt.Println("输入回车返回上一级")
		fmt.Scanf("%d\n", &x)
		break
	}
}
func selectuser() {
	dsn := "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}) //没有表格则创建一个user表格
	var ID, x int
	var N string
	user := User{}
	fmt.Println("1，输入ID\n2，输入名字\n3,查看全部用户\n4,返回上一级")
	fmt.Scanf("%d\n", &x)
	switch x {
	case 1:
		fmt.Println("请输入ID")
		fmt.Scanf("%d\n", &ID)
		result := db.First(&user, ID)
		if result != nil {
			fmt.Println(user.ID, user.Name, user.Gender)
			Return() //返回上一级
		} else {
			fmt.Println("查找不到该用户")
			Return() //返回上一级
		}

		time.Sleep(time.Second)
	case 2:
		users := []User{}
		fmt.Println("请输入姓名")
		fmt.Scanf("%s\n", &N)
		db.Where("name = ?", N).Find(&users)
		fmt.Println(users)
		Return() //返回上一级
		time.Sleep(time.Second)
	case 3:
		users := []User{}
		db.Find(&users)
		fmt.Println(users)
		Return() //返回上一级
	case 4:
		return
	default:
		fmt.Println("输入有误")
		Return() //返回上一级

	}
} //查找用户
func changeuser() {
	dsn := "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}) //没有表格则创建一个user表格
	var ID, x int
	var N, G string
	user := User{}
	fmt.Println("输入更改用户的ID")
	fmt.Scanf("%d\n", &ID)
	fmt.Println("1，更改名字\n2,更改性别\n3,返回上一级")
	fmt.Scanf("%d\n", &x)
	switch x {
	case 1:
		fmt.Println("请输入更改后的名字")
		fmt.Scanf("%s\n", &N)
		db.First(&user, ID)
		user.Name = N
		db.Save(&user)
		fmt.Println("更改成功")
		Return() //返回上一级
		time.Sleep(time.Second)
	case 2:
		fmt.Println("请输入更改后的性别")
		fmt.Scanf("%s\n", &G)
		db.First(&user, ID)
		user.Gender = G
		db.Save(&user)
		fmt.Println("更改成功")
		Return() //返回上一级
		time.Sleep(time.Second)
	case 3:
		return
	default:
		fmt.Println("输入有误")
		Return() //返回上一级

	}
} //更改用户

func main() {
	x := 0
	for {
		fmt.Println("欢迎使用数据库系统\n1，增加用户\n2，删除用户\n3，更改用户\n4，查找用户\n5,退出系统")
		fmt.Scanf("%d\n", &x)
		switch x {
		case 1:
			adduser()

		case 2:
			deleteuser()

		case 3:
			changeuser()

		case 4:
			selectuser()

		case 5:
			return
		default:
			fmt.Println("输入数字有误")
			time.Sleep(time.Second) //有误则菜单延迟一秒出现
		}

	}

}
