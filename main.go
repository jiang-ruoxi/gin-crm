package main

import (
	"gin-crm/model/mysql"
	"gin-crm/router"
	"gin-crm/lib"
	"log"
)

func main()  {
	serverConfig := lib.LoadServerConfig()
	mysql.InitDB(serverConfig)
	defer mysql.DB.Close()
	r := router.InitRouter()
	if err := r.Run(); err != nil {
		log.Fatal("服务器启动失败...")
	}
}
