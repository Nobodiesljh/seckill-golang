package routers

import (
	"github.com/gin-gonic/gin"
	"seckill-golang/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 商品信息展示页面获取数据
	r.GET("/good", controller.ShowGoods)
	return r
}
