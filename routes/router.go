package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

// InitRouter （可跨包） initRouter（私有方法，只能包内） 大小写问题
func InitRouter() {

	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		//http://localhost:3000/api/v1/hello

		// 用户模块的路由
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		// http://localhost:3000/api/v1//user/2  软删除
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("categorys", v1.GetCategorys)
		routerV1.PUT("category/:id", v1.EditCategory)
		routerV1.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由
		routerV1.POST("article/add", v1.AddArticles)
		routerV1.GET("articles", v1.GetArticles)
		routerV1.PUT("article/:id", v1.EditArticle)
		routerV1.DELETE("article/:id", v1.DeleteArticles)

		// 登录模块的路由

	}
	r.Run(utils.HttpPort)

}
