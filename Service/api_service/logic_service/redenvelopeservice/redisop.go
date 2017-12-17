package redenvelopeservice

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/lufred/red_envelope/Service/api_service/config"
)

var lockPrefix = "lock_"
var receiveGroupPrefix = "rv_"
var lockExpiration = 50 * time.Second

//Rede 红包redis存储结构
type Rede struct {
	ID         int32  `json:"id"`
	Amount     int32  `json:"amount"`
	Count      int32  `json:"count"`
	SecretCode string `json:"secret_code"`
	Expire     int64  `json:"expire"`
}

var charList = []string{
	`0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`,
	`8`, `9`, `a`, `b`, `c`, `d`, `e`, `f`,
	`g`, `h`, `i`, `j`, `k`, `l`, `m`, `n`,
	`o`, `p`, `q`, `r`, `s`, `t`, `u`, `v`,
	`w`, `x`, `y`, `z`,
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetSecretCode 随机获取8位口令码
func GetSecretCode() string {
	var code string
	for i := 0; i < 8; i++ {
		n := rand.Intn(36)
		code = code + charList[n]
	}
	return code
}

//Lock 获取红包锁
func Lock(id int32) bool {
	var curTime = time.Now()
	client := GetRedisClient()
	for (time.Now().UnixNano() - curTime.UnixNano()) < int64(lockExpiration) {
		if client.SetNX(lockPrefix+strconv.Itoa(int(id)), 1, lockExpiration).Val() {
			return true
		}
		time.Sleep(10 * time.Millisecond)
	}
	return false
}

//UnLock 释放红包锁
func UnLock(id int32) int64 {
	client := GetRedisClient()
	return client.Del(lockPrefix + strconv.Itoa(int(id))).Val()
}

//GetRedisClient 获取redis连接对象
func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.ProConfig.Redis.Addr,
		DB:   0, // use default DB
	})
	return client
}

//insertRedEToRedis 添加新红包到redis
func insertRedEToRedis(rede *Rede) error {
	client := GetRedisClient()
	value, _ := json.Marshal(rede)
	cmd := client.Set(strconv.Itoa(int(rede.ID)), value, time.Duration(rede.Expire*1e6-time.Now().UnixNano()))
	return cmd.Err()
}

//InsertUserToGroup 将用户插入到已领取队列
func InsertUserToGroup(rid int32, uid int32) (int64, error) {
	client := GetRedisClient()
	cmd := client.SAdd(receiveGroupPrefix+strconv.Itoa(int(rid)), uid)
	return cmd.Val(), cmd.Err()
}
