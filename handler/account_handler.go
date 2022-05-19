/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:23:25
 * @FilePath: \go-test\handler\account_handler.go
 * @Description: handler 层
 */
package handler

import (
	"gin-api/model"
	"gin-api/pkg/errno"
	"gin-api/service"

	"github.com/gin-gonic/gin"
)

/**
 * @description:Sign up API
 * @param undefined
 * @return {*}
 */
func Signup(c *gin.Context) {
	form := &model.SignupReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.InitAccountService()
	resp, responseErr := service.Signup(form)
	if responseErr != nil {
		SendResponse(c, responseErr, nil)
		return
	}
	SendResponse(c, nil, resp)
}

/**
 * @description: Sign in API
 * @param undefined
 * @return {*}
 */
func Signin(c *gin.Context) {
	form := &model.SigninReq{}
	if err := c.ShouldBindJSON(form); err != nil {
		SendResponse(c, errno.ParamError, nil)
		return
	}

	service := service.InitAccountService()

	resp, responseErr := service.Signin(form)
	if responseErr != nil {
		SendResponse(c, responseErr, nil)
		return
	}

	SendResponse(c, nil, resp)
}

/**
 * @description: Profile API
 * @param undefined
 * @return {*}
 */
func Profile(c *gin.Context) {
	service := service.InitAccountService()
	resp, responseErr := service.Profile()
	if responseErr != nil {
		SendResponse(c, responseErr, nil)
		return
	}
	SendResponse(c, nil, resp)
}

/**
 * @description:Update API
 * @param undefined
 * @return {*}
 */
func Update(c *gin.Context) {
	form := &model.UpdateReq{}
	if err := c.ShouldBindJSON(form); err != nil {
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.InitAccountService()
	resp, responseErr := service.Update(form)
	if responseErr != nil {
		SendResponse(c, responseErr, nil)
		return
	}
	SendResponse(c, nil, resp)
}
