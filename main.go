/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:41:59
 * @FilePath: \go-test\main.go
 * @Description: 程序入口类 (Program entry class)
 */
package main

import (
	"fmt"
	"gin-api/pkg"
	"gin-api/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	pkg.Init(*cfg)

	gin.SetMode(viper.GetString("run_mode"))

	g := gin.New()

	router.Load(
		g,
	)
	fmt.Printf("start to listening http address: %s", viper.GetString("addr"))
	fmt.Printf("ListenAndServe:%v", http.ListenAndServe(viper.GetString("addr"), g).Error())
}
