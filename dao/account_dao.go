/*
 * @Author: 呂健 10076418lv_jian@cn.tre-inc.com
 * @Date: 2022-05-18 15:05:02
 * @LastEditors: 呂健 10076418lv_jian@cn.tre-inc.com
 * @LastEditTime: 2022-05-18 17:26:19
 * @FilePath: \go-test\dao\account_dao.go
 * @Description: 操作数据库(control database)
 */
package dao

import (
	"fmt"
	"gin-api/model"
	"gin-api/pkg"

	"github.com/jmoiron/sqlx"
)

type AccountDao struct {
	db *sqlx.DB
}

func NewAccountDao() *AccountDao {
	return &AccountDao{pkg.DB}
}

/**
 * @description: Insert A Account
 * @param firstName
 * @param lastName
 * @param email
 * @param password
 * @return {*}
 */
func (dao *AccountDao) InsertAccount(firstName string, lastName string, email string, password string) error {
	sql := fmt.Sprintf(`insert into tm_account(FirstName,LastName,Email,Password) `)
	sql += fmt.Sprintf(`values (?,?,?,?) `)
	fmt.Printf("InsertAccount sql:%v", sql)

	rows, err := dao.db.Queryx(sql, firstName, lastName, email, password)
	if err != nil {
		fmt.Printf("InsertAccount error:%v", sql)
		return err
	}
	defer rows.Close()
	return nil
}

/**
 * @description: Get Account by Email
 * @param email
 * @return {*}
 */
func (dao *AccountDao) GetAccountByEmail(email string) (*model.ProfileResp, error) {
	sql := fmt.Sprintf(`select FirstName,LastName,Email `)
	sql += fmt.Sprintf(`from tm_account `)
	sql += fmt.Sprintf(`where Email=? `)
	fmt.Printf("GetAccountByEmail sql:%v", sql)

	rows, err := dao.db.Queryx(sql, email)
	if err != nil {
		fmt.Printf("GetAccountByEmail error:%v", sql)
		return nil, err
	}
	defer rows.Close()
	item := &model.ProfileResp{}
	for rows.Next() {
		rows.StructScan(item)
	}
	return item, nil
}

/**
 * @description: Get Account by Email and Password
 * @param email
 * @param password
 * @return {*}
 */
func (dao *AccountDao) GetAccountByEmailAndPassword(email string, password string) (*model.ProfileResp, error) {
	sql := fmt.Sprintf(`select FirstName,LastName,Email `)
	sql += fmt.Sprintf(`from tm_account `)
	sql += fmt.Sprintf(`where Email=? `)
	sql += fmt.Sprintf(`and Password=? `)
	fmt.Printf("GetAccountByEmailAndPassword sql:%v", sql)

	rows, err := dao.db.Queryx(sql, email, password)
	if err != nil {
		fmt.Printf("GetAccountByEmailAndPassword error:%v", sql)
		return nil, err
	}
	defer rows.Close()
	item := &model.ProfileResp{}
	for rows.Next() {
		rows.StructScan(item)
	}
	return item, nil
}

/**
 * @description: Update Account by Email and Password
 * @param email
 * @param password
 * @param firstName
 * @param lastName
 * @return {*}
 */
func (dao *AccountDao) UpdateAccountByEmailAndPassword(email string, password string, firstName string, lastName string) error {
	sql := fmt.Sprintf(`update tm_account set FirstName=?,LastName=? `)
	sql += fmt.Sprintf(`where Email=? `)
	sql += fmt.Sprintf(`and Password=? `)
	fmt.Printf("UpdateAccountByEmailAndPassword sql:%v", sql)

	rows, err := dao.db.Queryx(sql, firstName, lastName, email, password)
	if err != nil {
		fmt.Printf("UpdateAccountByEmailAndPassword error:%v", sql)
		return err
	}
	defer rows.Close()
	return nil
}
