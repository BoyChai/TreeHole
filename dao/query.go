package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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
func (d *dao) GetUser(id string) (User, error) {
	var u User
	if err := d.db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户未找到，返回自定义错误
			return User{}, err
		}
		// 打印详细的错误信息
		fmt.Printf("查询用户时发生错误: %v\n", err)
		return User{}, errors.New("查询用户时发生错误")
	}
	return u.Clear(), nil
}

// GetArticles 获取帖子列表
func (d *dao) GetArticles() ([]Article, error) {
	var articles []Article
	tx := d.db.Find(&articles, "deleted_at IS NULL")
	if tx.Error != nil {
		return []Article{}, errors.New(fmt.Sprintf("帖子查询错误: %s", tx.Error.Error()))
	}
	return articles, nil
}

// GetArticle 获取帖子内容
func (d *dao) GetArticle(id string) (Article, error) {
	var article Article
	tx := d.db.First(&article, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			// 记录未找到，返回自定义错误
			return Article{}, tx.Error
		}
		return Article{}, errors.New(fmt.Sprintf("帖子查询错误: %s", tx.Error.Error()))
	}
	return article, nil
}

// GetHomeComment 获取10个评论内容
func (d *dao) GetHomeComment() ([]Comment, error) {
	var comments []Comment

	// 查询最新的10条评论，且这些评论不允许有父级评论
	result := d.db.Order("created_at desc").Limit(10).Where("parent = ?", 0).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

// GetComment 根据帖子获取评论列表
func (d *dao) GetComment(id string) ([]Comment, error) {
	var comments []Comment
	tx := d.db.Model(Comment{}).Where("article = ?", id).Find(&comments)
	if tx.Error != nil {
		return []Comment{}, errors.New(fmt.Sprintf("帖子查询错误: %s", tx.Error))
	}
	return comments, nil
}
