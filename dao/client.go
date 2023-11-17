package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Client 数据库链接
func Client() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user",
		"password",
		"localhost",
		3306,
		"databases",
	)
	Dao.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键（代码里自动体验外键关系）
	})
	if err != nil {
		fmt.Println("链接数据库错误:", err)
		return
	}
}

func AutoTables() {
	var err error
	// 用户表
	err = Dao.db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
	// 帖子表
	err = Dao.db.AutoMigrate(&Article{})
	if err != nil {
		fmt.Println(err)
	}
	// 评论表
	err = Dao.db.AutoMigrate(&Comment{})
	if err != nil {
		fmt.Println(err)
	}
}
