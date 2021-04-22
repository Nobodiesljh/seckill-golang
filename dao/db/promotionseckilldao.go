package db

import "seckill-golang/models"

func UpdateCountByGoodsId(goodsId int64) error {
	sqlStr := "update t_promotion_seckill set ps_count = 100, version = 0 where goods_id = ?"
	_, err := DB.Exec(sqlStr, goodsId)
	if err != nil {
		return err
	}
	return nil
}

func SelectCountByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "select ps_count from t_promotion_seckill where goods_id = ?"
	var count int64
	err := DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}

func SelectCountByGoodsIdPcc(goodsId int64) (int64, error) {
	sqlStr := "select ps_count from t_promotion_seckill where goods_id = ? for update"
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

func ReduceByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "update t_promotion_seckill set ps_count = ps_count-1 where ps_count>0 and goods_id = ?"
	res, err := DB.Exec(sqlStr, goodsId)
	count, getErr := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if getErr != nil {
		return 0, getErr
	}
	return count, nil
}

func SelectGoodByGoodsId(goodsId int64) (models.PromotionSecKill, error) {
	sqlStr := "select * from t_promotion_seckill where goods_id = ?"
	var promotion models.PromotionSecKill
	err := DB.Get(&promotion, sqlStr, goodsId)
	if err != nil {
		return promotion, err
	}
	return promotion, nil
}

func ReduceStockByOcc(goodsId int64, num int64, version int64) (int64, error) {
	sqlStr := "update t_promotion_seckill set ps_count = ps_count-?, version = version+1 " +
		"where version = ? and goods_id = ?"
	res, err := DB.Exec(sqlStr, num, version, goodsId)
	count, getErr := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if getErr != nil {
		return 0, getErr
	}
	return count, nil
}
