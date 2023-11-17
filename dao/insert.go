package dao

import (
	"errors"
	"fmt"
	"time"
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

// CreateArticle 创建帖子
func (d *dao) CreateArticle(user uint, text string, title string) (Article, error) {
	order := Article{
		Sender:    user,
		Text:      text,
		Title:     title,
		CreatedAt: time.Now(),
	}
	tx := d.db.Create(&order)
	if tx.Error != nil {
		return Article{}, errors.New(fmt.Sprint("创建帖子出现错误:", tx.Error.Error()))
	}
	return order, nil
}

// SendComment 发送评论
func (d *dao) SendComment(user uint, article uint, text string, parent uint) (Comment, error) {
	comment := Comment{
		Sender:  user,
		Article: article,
		Text:    text,
		Parent:  parent,
	}
	tx := d.db.Create(&comment)
	if tx.Error != nil {
		return Comment{}, errors.New(fmt.Sprint("创建帖子出现错误:", tx.Error.Error()))
	}
	return comment, nil
}
