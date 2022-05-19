/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:28:29
 * @FilePath: \go-test\model\account_model.go
 * @Description: Struct class
 */
package model

type SignupReq struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type SignupResp struct {
	FirstName string
	LastName  string
	Email     string
}

type SigninReq struct {
	Email    string
	Password string
}

type SigninResp struct {
	Token string
}

type ProfileResp struct {
	FirstName string `db:"FirstName" json:"FirstName"`
	LastName  string `db:"LastName" json:"LastName"`
	Email     string `db:"Email" json:"Email"`
}

type UpdateReq struct {
	FirstName string
	LastName  string
}
