package utils

import (
	"math"
	"sync"
	"time"
)

const (
	epoch  = uint64(1642780800000) // 设置起始毫秒：2022-01-13 00:00:00
	bitSeq = uint64(8)             // 序列所占的位数
	maxSeq = uint64(math.MaxUint8) // 支持的最大序列id数量
)

type snowFlake struct {
	seq   uint64
	time  uint64
	mutex sync.Mutex
}

var sf *snowFlake = &snowFlake{}

func GetVal() uint64 {
	sf.mutex.Lock()
	mill := uint64(time.Now().UnixMilli())
	if mill != sf.seq || sf.seq >= maxSeq {
		sf.seq = 0
		sf.time = mill
	}
	sf.seq++
	ret := ((sf.time - epoch) << bitSeq) + sf.seq
	sf.mutex.Unlock()
	return ret
}
