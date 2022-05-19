/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:29:23
 * @FilePath: \go-test\pkg\errno\code.go
 * @Description: Response Code
 */
package errno

var (
	SUCCESS                 = &Errno{Code: 200, Msg: "success"}
	ParamError              = &Errno{Code: 400, Msg: "Parameter validation error"}
	NameValidationError     = &Errno{Code: 400, Msg: "[First name and last name validation error] the first name and last name cannot be empty and more than 64 characters"}
	EmailValidationError    = &Errno{Code: 400, Msg: "[Email address validation error] the email address has to be unique"}
	PasswordValidationError = &Errno{Code: 400, Msg: "[Password validation error]the password has to be minimum 6 characters,maximum 16 characters and alphanumeric"}
	InternalError           = &Errno{Code: 500, Msg: "database or api internal panic"}
)
