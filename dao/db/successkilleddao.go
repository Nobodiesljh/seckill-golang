package db

import (
	"seckill-golang/models"
)

func DeleteByGoodsId(goodsId int64) error {
	sqlStr := "delete from t_success_killed where goods_id = ?"
	_, err := DB.Exec(sqlStr, goodsId)
	if err != nil {
		return err
	}
	return nil
}

func CreatOrder(killed models.SuccessKilled) error {
	sqlStr := "insert into t_success_killed (goods_id, user_id, state, create_time) values (?, ?, ?, ?)"
	_, err := DB.Exec(sqlStr, killed.GoodsId, killed.UserId, killed.State, killed.CreateTime)
	if err != nil {
		return err
	}
	return nil
}

func GetKilledCountByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "select count(*) from t_success_killed where goods_id = ?"
	var count int64
	err := DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}
