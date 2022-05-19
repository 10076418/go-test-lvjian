/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:32:15
 * @FilePath: \go-test\service\account_service.go
 * @Description: Service Class 业务逻辑层
 */
package service

import (
	"fmt"
	"gin-api/dao"
	"gin-api/model"
	"gin-api/pkg"
	"gin-api/pkg/errno"
	"gin-api/util"
)

type AccountService struct {
	accountDao *dao.AccountDao
}

func InitAccountService() *AccountService {
	return &AccountService{dao.NewAccountDao()}
}

func (service *AccountService) Signup(form *model.SignupReq) (*model.SignupResp, *errno.Errno) {
	fmt.Printf("Signup Begin:%v", form)

	//Email validation
	verifyEmailFlag := util.VerifyEmailFormat(form.Email)
	if verifyEmailFlag == false {
		return nil, errno.EmailValidationError
	}
	pwdCrypt := util.MD5(form.Password)
	account, dataErr := service.accountDao.GetAccountByEmail(form.Email)
	if dataErr != nil {
		return nil, errno.InternalError
	}
	if account.Email != "" {
		return nil, errno.EmailValidationError
	}
	//FirstName validation
	verifyFirstNameFlag := util.VerifyNameFormat(form.FirstName)
	if verifyFirstNameFlag == false {
		return nil, errno.NameValidationError
	}
	//LastName validation
	verifyLastNameFlag := util.VerifyNameFormat(form.LastName)
	if verifyLastNameFlag == false {
		return nil, errno.NameValidationError
	}
	//Password validation
	verifyPasswordFlag := util.VerifyPasswordFormat(form.Password)
	if verifyPasswordFlag == false {
		return nil, errno.PasswordValidationError
	}
	insertErr := service.accountDao.InsertAccount(form.FirstName, form.LastName, form.Email, pwdCrypt)
	if insertErr != nil {
		return nil, errno.InternalError
	}
	resp := &model.SignupResp{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
	}
	fmt.Printf("Signup End:%v", resp)
	return resp, nil
}

func (service *AccountService) Signin(form *model.SigninReq) (*model.SigninResp, *errno.Errno) {
	fmt.Printf("Signin Begin:%v", form)
	//Email validation
	verifyEmailFlag := util.VerifyEmailFormat(form.Email)
	if verifyEmailFlag == false {
		return nil, errno.EmailValidationError
	}
	//Password validation
	verifyPasswordFlag := util.VerifyPasswordFormat(form.Password)
	if verifyPasswordFlag == false {
		return nil, errno.PasswordValidationError
	}
	pwdCrypt := util.MD5(form.Password)
	account, dataErr := service.accountDao.GetAccountByEmailAndPassword(form.Email, pwdCrypt)
	if dataErr != nil {
		return nil, errno.InternalError
	}
	if account.Email == "" {
		return nil, nil
	}
	token, tokenErr := signApiToken(form.Email, form.Password)
	if tokenErr != nil {
		return nil, errno.InternalError
	}
	resp := &model.SigninResp{
		Token: token,
	}
	fmt.Printf("Signin End:%v", resp)
	key := fmt.Sprint(`redis_token`)
	if err := pkg.SetRedisVal(key, resp, 0); err != nil {
		return nil, errno.InternalError
	}
	return resp, nil
}

func (service *AccountService) Profile() (*model.ProfileResp, *errno.Errno) {
	fmt.Printf("Profile Begin")
	pwdCrypt := util.MD5(pkg.StaticPassword)

	account, err := service.accountDao.GetAccountByEmailAndPassword(pkg.StaticEmail, pwdCrypt)
	if err != nil {
		return nil, errno.InternalError
	}
	if account.Email == "" {
		return nil, nil
	}
	resp := &model.ProfileResp{
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
	fmt.Printf("Profile End:%v", resp)
	return resp, nil
}

func (service *AccountService) Update(form *model.UpdateReq) (*model.ProfileResp, *errno.Errno) {
	fmt.Printf("Update Begin")
	//FirstName validation
	verifyFirstNameFlag := util.VerifyNameFormat(form.FirstName)
	if verifyFirstNameFlag == false {
		return nil, errno.NameValidationError
	}
	//LastName validation
	verifyLastNameFlag := util.VerifyNameFormat(form.LastName)
	if verifyLastNameFlag == false {
		return nil, errno.NameValidationError
	}
	pwdCrypt := util.MD5(pkg.StaticPassword)

	err := service.accountDao.UpdateAccountByEmailAndPassword(pkg.StaticEmail, pwdCrypt, form.FirstName, form.LastName)
	if err != nil {
		return nil, errno.InternalError
	}
	account, dataErr := service.accountDao.GetAccountByEmailAndPassword(pkg.StaticEmail, pwdCrypt)
	if dataErr != nil {
		return nil, errno.InternalError
	}
	if account.Email == "" {
		return nil, nil
	}
	resp := &model.ProfileResp{
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
	fmt.Printf("Update End:%v", resp)
	return resp, nil
}

func signApiToken(email string, password string) (string, error) {
	fmt.Printf("signApiToken Begin:%v %v", email, password)

	token, err := pkg.Sign(pkg.Token{Email: email, Password: password})
	if err != nil {
		return "", nil
	}
	fmt.Printf("signApiToken End:%v", token)
	return token, nil
}
