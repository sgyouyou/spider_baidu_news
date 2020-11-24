package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	News      News
	connectDb func() gorm.DB
	insertNew func(news News)
}

/**
 * 文章表结构
 */
type News struct {
	Title   string `gorm:"column:title"` // 定义字段名和对应在表中的字段名
	Author  string `gorm:"column:author"`
	Time    int64  `gorm:"column:create_time"`
	Content string `gorm:"column:content"`
}

/**
 * 连接数据库
 */
func (p Post) ConnectDb() gorm.DB {
	dsn := "walter:123456@tcp(192.168.31.254:3306)/lara?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return *db
}

/**
 * 插入数据
 */
func (p Post) InsertNew(row News) {
	db := p.ConnectDb()
	res := db.Create(row)
	fmt.Printf("Error: %v, RowsAffected: %v\n", res.Error, res.RowsAffected)
}
