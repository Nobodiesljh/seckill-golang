package common

import "sync"

type singleton chan [2]int64

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		// 容量视库存大小而定，这里设为100
		ret := make(singleton, 100)
		instance = &ret
	})
	return instance
}
