/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:30:11
 * @FilePath: \go-test\pkg\api_token.go
 * @Description:Token Class
 */
package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrMissingHeader = errors.New("the length of the authorization header is zero")
	StaticEmail      string
	StaticPassword   string
)

type Token struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func ParseRequest(c *gin.Context) (*Token, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return &Token{}, ErrMissingHeader
	}

	secret := viper.GetString("trechina_jwt")

	var t string
	fmt.Sscanf(header, "%s", &t)
	return Parse(t, secret)
}

func Parse(tokenString string, secret string) (*Token, error) {
	ctx := &Token{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Email = claims["Email"].(string)
		StaticEmail = ctx.Email
		ctx.Password = claims["Password"].(string)
		StaticPassword = ctx.Password
		return ctx, nil
	} else {
		return ctx, err
	}
}

func Sign(c Token) (tokenString string, err error) {
	secret := viper.GetString("trechina_jwt")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":    c.Email,
		"Password": c.Password,
		"exp":      time.Now().Unix() + 60*10,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
