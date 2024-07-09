package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))

	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	// 自动迁移 表名user和struct User之间的映射关系
	db.AutoMigrate(&User{}, &Category{}, &Article{})
	// 禁用表名的复数形式 比如 User ->users
	db.SingularTable(true)

	// 连接池设计
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second) //注意不要大于矿建里的时间

	//db.Close()

}
