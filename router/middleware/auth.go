/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:31:32
 * @FilePath: \go-test\router\middleware\auth.go
 * @Description: JWT Class
 */
package middleware

import (
	"fmt"
	"gin-api/handler"
	"gin-api/pkg"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := pkg.ParseRequest(c)
		if err != nil {
			fmt.Printf("AuthMiddleware error:%v", err)
			handler.SendResponse(c, err, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
