/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:32:45
 * @FilePath: \go-test\util\common_util.go
 * @Description: 工具类 Tool Class
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
)

/*
MD5 encode
*/
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

/*
Verify Email
*/
func VerifyEmailFormat(email string) bool {
	fmt.Printf("VerifyEmailFormat:%v", email)

	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/*
Verify FirstName or LastName
*/
func VerifyNameFormat(name string) bool {
	charray := []byte(name)
	fmt.Printf("VerifyNameFormat:%v", len(charray))

	if len(charray) > 64 || name == "" {
		return false
	}
	return true
}

/*
Verify Password
*/
func VerifyPasswordFormat(password string) bool {
	fmt.Printf("VerifyPasswordFormat:%v", password)
	re, err1 := regexp.Compile(`^[0-9A-Za-z]{5,16}$`)
	if err1 != nil {
		return false
	}
	match1 := re.MatchString(password)
	if !match1 {
		return false
	}
	re, err2 := regexp.Compile(`^[0-9]{5,16}$`)
	if err2 != nil {
		return false
	}
	match2 := re.MatchString(password)
	if match2 {
		return false
	}
	re, err3 := regexp.Compile(`^[A-Za-z]{5,16}$`)
	if err3 != nil {
		return false
	}
	match3 := re.MatchString(password)
	if match3 {
		return false
	}
	return true
}
