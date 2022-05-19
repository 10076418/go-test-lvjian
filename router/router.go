/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:54:07
 * @FilePath: \go-test\router\router.go
 * @Description: Rounter Class
 */
package router

import (
	"gin-api/handler"
	"gin-api/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: Init API
 */
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect api route")
	})

	account := g.Group("/v1")
	{
		account.POST("/signup", handler.Signup)
		account.POST("/signin", handler.Signin)
	}

	userAuth := g.Group("/v1/profile/")
	userAuth.Use(middleware.AuthMiddleware())
	{
		userAuth.GET("/", handler.Profile)
		userAuth.POST("/update", handler.Update)
	}

	return g
}
