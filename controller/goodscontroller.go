package controller

import (
	"github.com/gin-gonic/gin"
	"seckill-golang/service"
	"strconv"
)

func ShowGoods(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)

	// 获取商品的 基本、封面、详情、参数信息
	goods := service.GetGoods(int64(id))
	covers := service.FindCovers(int64(id))
	details := service.FindDetails(int64(id))
	params := service.FindParams(int64(id))

	// 往前端返回查询到的数据
	c.JSON(200, gin.H{
		"goods":   goods,
		"covers":  covers,
		"details": details,
		"params":  params,
	})
}
