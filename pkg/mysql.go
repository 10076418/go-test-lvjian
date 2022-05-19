/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:30:57
 * @FilePath: \go-test\pkg\mysql.go
 * @Description: Mysql Class
 */
package pkg

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

func LoadDbConfig(viper *viper.Viper) *DBConfig {
	cfg := &DBConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	fmt.Printf("LoadDbConfig:%+v\n", cfg)
	return cfg
}

func (cfg *DBConfig) InitDB() *sqlx.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("InitDb error:%v", dataSourceName)
		panic("InitDb error")
	}
	return db
}
