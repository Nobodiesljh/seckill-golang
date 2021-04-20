package main

import (
	"seckill-golang/dao/db"
	"seckill-golang/routers"
)

func main() {
	// 数据库初始化
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routers.SetupRouter()
	r.Run("127.0.0.1:8080")
}
