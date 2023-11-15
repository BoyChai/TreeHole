package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

var Dao dao

type dao struct {
	db *gorm.DB
}
type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan is used to convert a database value to the custom JSON type.
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid JSON format")
	}
	return json.Unmarshal(bytes, j)
}

// User 用户表
type User struct {
	// ID用户的唯一标识
	ID uint `gorm:"primaryKey;auto_increment"`
	// 用户名
	Name string `gorm:"not null"`
	// 电话号
	Number string `gorm:"unique;not null"`
	// 密码(md5(Pass))
	Pass string `gorm:"not null"`
	// 图片地址
	ImgUrl string `gorm:"not null"`
	// 邮箱
	Email string `gorm:"unique"`
	// 角色
	Role Role `gorm:"not null"`
	// 其他数据
	Other JSON
}
type Role int

// Article 帖子列表
type Article struct {
	// 帖子唯一id
	ID uint `gorm:"primaryKey;auto_increment"`
	// 帖子创建时间
	CreatedAt time.Time
	// 帖子更新时间
	UpdatedAt time.Time
	// 帖子删除时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// 帖子标题
	Title string
	// 帖子内容
	Text string `gorm:"type: text"`
	// 发布者
	Sender uint
	// 修改者
	Modified uint
}

type Comment struct {
	// 评论的唯一id
	ID uint `gorm:"primaryKey;auto_increment"`
	// 评论创建时间
	CreatedAt time.Time
	// 帖子删除时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// 评论发布者id
	Sender uint
	// 帖子id
	Article uint
	// 发送内容
	Text string `gorm:"type: text"`
	// 父级评论
	Parent uint
}
