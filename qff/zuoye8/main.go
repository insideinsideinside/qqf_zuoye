package main

import (
	"main/boot"
)

func main() {
	boot.ServerSetup()
	boot.ViperSetup("./manifest/config/config.yaml")
	boot.Loggersetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
}
