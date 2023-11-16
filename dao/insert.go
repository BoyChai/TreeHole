package dao

import (
	"errors"
	"fmt"
)

// CreateUser 创建用户
func (d *dao) CreateUser(name string, number string, pass string, imgUrl string, mail string, role Role) error {
	user := User{
		Name:   name,
		Number: number,
		Pass:   pass,
		ImgUrl: imgUrl,
		Email:  mail,
		Role:   role,
		Other:  JSON{},
	}
	tx := d.db.Create(&user)
	if tx.Error != nil {
		return errors.New(fmt.Sprint("创建用户出现错误:", tx.Error.Error()))
	}
	return nil
}
