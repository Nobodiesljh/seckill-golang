package service

import (
	"fmt"
	"seckill-golang/dao/db"
	"seckill-golang/models"
	"sync"
	"time"
)

var lock sync.Mutex

func InitializeSecKill(goodsId int64) {
	tx, err := db.DB.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
	}

	err1 := db.DeleteByGoodsId(goodsId)
	err2 := db.UpdateCountByGoodsId(goodsId)

	if err1 != nil {
		fmt.Println(err1)
		tx.Rollback()
	}
	if err2 != nil {
		fmt.Println(err2)
		tx.Rollback()
	}

	// 提交事务
	tx.Commit()
}

func HandleSeckill(goodsId int64, userId int64) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	// 检查库存
	count, errCount := db.SelectCountByGoodsId(goodsId)
	if errCount != nil {
		return errCount
	}

	if count > 0 {
		// 1.扣库存
		errRed := db.ReduceStockByGoodsId(goodsId, count-1)
		if errRed != nil {
			tx.Rollback()
			return errRed
		}

		// 2.创建订单
		killed := models.SuccessKilled{
			GoodsId:    goodsId,
			UserId:     userId,
			State:      0,
			CreateTime: time.Now(),
		}
		errCre := db.CreatOrder(killed)
		if errCre != nil {
			tx.Rollback()
			return errCre
		}
	}
	tx.Commit()
	return nil
}

func HandleSecKillWithLock(goodsId int64, userId int64) error {
	lock.Lock()
	err := HandleSeckill(goodsId, userId)
	lock.Unlock()
	return err
}

func GetKilledCount(goodsId int64) (int64, error) {
	return db.GetKilledCountByGoodsId(goodsId)
}
