/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:31:11
 * @FilePath: \go-test\pkg\redis.go
 * @Description: Redis Class
 */
package pkg

import (
	"encoding/json"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
)

var (
	MaxErrCnt = 5
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

func LoadRedisConfig(viper *viper.Viper) *RedisConfig {
	cfg := &RedisConfig{
		Addr:     viper.GetString("host"),
		Password: viper.GetString("password"),
		DB:       viper.GetInt("index"),
		PoolSize: viper.GetInt("pool_size"),
	}
	return cfg
}

func (cfg *RedisConfig) InitRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	if redisCli == nil {
		panic("redis.NewClient error")
	}
	_, err := redisCli.Ping().Result()
	if err != nil {
		panic("redisCli.Ping error")
	}
	return redisCli
}

func SetRedisVal(key string, val interface{}, expiration time.Duration) error {
	buf, err := json.Marshal(val)
	if err != nil {
		return err
	}
	cmd := RedisCli.Set(key, string(buf), expiration)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func GetRedisVal(key string) string {
	tryCnt := 0
	for {
		value, err := RedisCli.Get(key).Result()
		if err == redis.Nil {
			return ""
		} else if err != nil {
			tryCnt++
			if tryCnt > MaxErrCnt {
				return ""
			}
		}
		return value
	}
}

func RedisExists(key string) bool {
	errCnt := 0
	for {
		ret := RedisCli.Exists(key)
		if ret == nil || ret.Err() != nil {
			errCnt++
		} else {
			return ret.Val()
		}
		if errCnt > MaxErrCnt {
			break
		}
	}
	return false
}
