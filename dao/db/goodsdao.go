package db

import (
	"fmt"
	"seckill-golang/models"
)

func FindGoodsById(goodsId int64) models.Goods {
	sqlStr := "select * from t_goods where goods_id = ?"
	var goods models.Goods
	err := DB.Get(&goods, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goods
}

func FindCoverByGoodsId(goodsId int64) []models.GoodsCover {
	sqlStr := "select * from t_goods_cover where goods_id = ? order by gc_order"
	var goodsCover []models.GoodsCover
	err := DB.Select(&goodsCover, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goodsCover
}

func FindDetailsByGoodsId(goodsId int64) []models.GoodsDetail {
	sqlStr := "select * from t_goods_detail where goods_id = ? order by gd_order"
	var goodsDetails []models.GoodsDetail
	err := DB.Select(&goodsDetails, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goodsDetails
}

func FindParamsByGoodsId(goodsId int64) []models.GoodsParam {
	sqlStr := "select * from t_goods_param where goods_id = ? order by gp_order"
	var goodsParams []models.GoodsParam
	err := DB.Select(&goodsParams, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goodsParams
}
