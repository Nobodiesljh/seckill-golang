package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seckill-golang/service"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func Handle(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)

	seckillNum := 50
	wg.Add(seckillNum)

	// 数据库中的商品、秒杀信息的初始化
	service.InitializeSecKill(int64(id))

	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := service.HandleSeckill(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	killedCount, err := service.GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "failed",
		"message": "会存在超卖现象",
	})
}

func HandleWithLock(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)

	seckillNum := 56
	wg.Add(seckillNum)

	// 数据库中的商品、秒杀信息的初始化
	service.InitializeSecKill(int64(id))

	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := service.HandleSecKillWithLock(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	killedCount, err := service.GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "success",
		"message": "秒杀正常",
	})
}
