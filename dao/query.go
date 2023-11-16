package dao

import (
	"errors"
	"fmt"
)

// CheckUser 查询用户是否合法,并返回用户数据
func (d *dao) CheckUser(number string, pass string) (User, error) {
	var u User
	tx := d.db.Where("number = ? and pass = ?", number, pass).First(&u)
	if tx.Error != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	u.Pass = ""
	return u, nil
}

// CheckNumberUser 通过手机号查询用户是否存在并返回用户数据
func (d *dao) CheckNumberUser(number string) (User, error) {
	var u User
	tx := d.db.Where("number = ?", number).First(&u)
	if tx.Error != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	u.Pass = ""
	return u, nil
}

// GetUser 通过id获取某个用户的信息
func (d *dao) GetUser(id uint) (User, error) {
	var u User
	tx := d.db.Where("id = ?", id).First(&u)
	if tx != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	return u.Clear(), nil
}
