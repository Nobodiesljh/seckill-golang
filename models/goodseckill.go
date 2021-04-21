package models

import (
	"time"
)

type PromotionSecKill struct {
	PsId         int64     `db:"ps_id"`
	GoodsId      int64     `db:"goods_id"`
	PsCount      int64     `db:"ps_count"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
	Status       int32     `db:"status"`
	CurrentPrice float64   `db:"current_price"`
	Version      int64     `db:"version"`
}

type SuccessKilled struct {
	GoodsId    int64     `db:"goods_id"`
	UserId     int64     `db:"user_id"`
	State      int16     `db:"state"`
	CreateTime time.Time `db:"create_time"`
}
