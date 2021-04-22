package routers

import (
	"github.com/gin-gonic/gin"
	"seckill-golang/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 商品信息展示页面获取数据
	r.GET("/good", controller.ShowGoods)

	// 单机锁
	seckillGroup := r.Group("/seckill")
	{
		// case1:不加锁,出现超卖现象
		seckillGroup.GET("/handle", controller.Handle)
		// case2:使用sync包中的Mutex类型的互斥锁,秒杀正常
		seckillGroup.GET("/handleWithLock", controller.HandleWithLock)
		// case4:数据库悲观锁(查询加锁),不能
		seckillGroup.GET("/handleWithPccOne", controller.HandleWithPccOne)
		// case5:数据库悲观锁(更新加锁),正常
		seckillGroup.GET("/handleWithPccTwo", controller.HandleWithPccTwo)
		// case6:数据库乐观锁，正常
		seckillGroup.GET("/handleWithOcc", controller.HandleWithOcc)
		// case7:GoLang中的channel，正常
		seckillGroup.GET("/handleWithChannel", controller.HandleWithChannel)
	}
	return r
}
