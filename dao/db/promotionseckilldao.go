package db

func UpdateCountByGoodsId(goodsId int64) error {
	sqlStr := "update t_promotion_seckill set ps_count = 100, version = 0 where goods_id = ?"
	_, err := DB.Exec(sqlStr, goodsId)
	if err != nil {
		return err
	}
	return nil
}

func SelectCountByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "select ps_count from t_promotion_seckill where goods_id = ?;"
	var count int64
	err := DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}

func ReduceStockByGoodsId(goodsId int64, count int64) error {
	sqlStr := "update t_promotion_seckill set ps_count = ? where goods_id = ?"
	_, err := DB.Exec(sqlStr, count, goodsId)
	if err != nil {
		return err
	}
	return nil
}
