package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	// 通过go向redis写入数据 string [key - value]
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 关闭连接
	defer conn.Close()

	// 读取数据 获取名字
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//nameString := r.(string)
	fmt.Println("Manipulate success, the name is", r)
} //不知道做的对不对。。
