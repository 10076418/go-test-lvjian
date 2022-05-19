/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:30:41
 * @FilePath: \go-test\pkg\config.go
 * @Description: Config Class
 */
package pkg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
)

var (
	Cfg      *Config
	DB       *sqlx.DB
	RedisCli *redis.Client
)

type Config struct {
	DB    *DBConfig
	Redis *RedisConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("viper.ReadInConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		DB:    LoadDbConfig(viper.Sub("db")),
		Redis: LoadRedisConfig(viper.Sub("redis")),
	}

	return cfg
}

func initConfig(cfg *Config) {
	DB = cfg.DB.InitDB()
	RedisCli = cfg.Redis.InitRedis()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed:%s", e.Name)
	})
}
