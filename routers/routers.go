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
	}
	return r
}
