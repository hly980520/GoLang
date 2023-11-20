package main

import (
	"fmt"
	"hs-admin-system-ksa/main/config"
	"hs-admin-system-ksa/main/router"
)

func main() {

	fmt.Println("Hello World")
	config.InitConfig()
	err := router.InitRouters()
	if err != nil {
		fmt.Println("启动失败 - ", err.Error())
		return
	}
	fmt.Println("启动成功")
}
