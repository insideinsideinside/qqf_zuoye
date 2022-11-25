package dao

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var database = map[string]string{}

func Readdatabase() { //读取并将信息导入假数据库
	database1, err1 := os.Open("username.txt")
	database2, err2 := os.Open("password.txt")
	if err1 != nil || err2 != nil {
		fmt.Println("Open file Failed", err1)
		fmt.Println("Open file Failed", err1)
		return
	}
	defer func() {
		database1.Close()
		database2.Close()
	}()
	reader1 := bufio.NewReader(database1) //获取读句柄
	reader2 := bufio.NewReader(database2) //获取读句柄
	for {
		line1, _, err := reader1.ReadLine()
		line2, _, err := reader2.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		database[string(line1)] = string(line2)
	}
}

func Writedatabase(username, password string) {
	database1, err := os.OpenFile("username.txt", os.O_APPEND|os.O_RDWR, 0644)
	database2, err := os.OpenFile("password.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open file Failed", err)
		return
	}
	defer func() {
		database1.Close()
		database2.Close()
	}()
	_, err = io.WriteString(database1, "\n")
	_, err = io.WriteString(database1, username)
	_, err = io.WriteString(database2, "\n")
	_, err = io.WriteString(database2, password)
	if err != nil {
		panic(err)
		return
	}

}
func Rewritedatabase(username, newpassword string) {

	database1, err := os.OpenFile("username.txt", os.O_APPEND|os.O_RDWR, 0644)
	database2, err := os.OpenFile("password.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open file Failed", err)
		return
	}
	defer func() {
		database1.Close()
		database2.Close()
	}()
	_, err = io.WriteString(database1, "\n")
	_, err = io.WriteString(database1, username)
	_, err = io.WriteString(database2, "\n")
	_, err = io.WriteString(database2, newpassword)
	if err != nil {
		panic(err)
		return
	}

}
func WriteMessage(username, message string) {
	database, err := os.OpenFile("message.txt", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Open file Failed", err)
		return
	}
	defer func() {
		database.Close()

	}()
	_, err = io.WriteString(database, "\n")
	_, err = io.WriteString(database, username)
	_, err = io.WriteString(database, "\n")
	_, err = io.WriteString(database, message)
	if err != nil {
		panic(err)
		return
	}

}

func AddUser(username, password string) {
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPassword(username string) string {
	return database[username]
}
func ChangePassword(username, password, newpassword string) {
	database[username] = newpassword
}
func FindPassword(username string) string {
	return database[username]
}
