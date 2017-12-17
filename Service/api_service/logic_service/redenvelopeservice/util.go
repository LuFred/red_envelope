package redenvelopeservice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetRandAmount 获取随机红包
func GetRandAmount(amount int32, count int32) int32 {
	if count == 1 {
		return amount
	}
	bal := amount - count
	return rand.Int31n(bal) + 1
}
