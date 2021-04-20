package service

import (
	"seckill-golang/dao/db"
	"seckill-golang/models"
)

func GetGoods(gid int64) models.Goods {
	return db.FindGoodsById(gid)
}

func FindCovers(gid int64) []models.GoodsCover {
	return db.FindCoverByGoodsId(gid)
}

func FindDetails(gid int64) []models.GoodsDetail {
	return db.FindDetailsByGoodsId(gid)
}

func FindParams(gid int64) []models.GoodsParam {
	return db.FindParamsByGoodsId(gid)
}
