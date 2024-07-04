package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {

	// 1.引用数据库
	model.InitDb()

	// 2.引用路由
	routes.InitRouter()
}
