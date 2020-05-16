package main

import (
	"github.com/SnDragon/gin_restful_demo/dao"
	"github.com/SnDragon/gin_restful_demo/model"
)

// 迁移工具,自动创建表
func main() {
	db := dao.NewDB()
	db.AutoMigrate(&model.User{})
}

