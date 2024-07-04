package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter （可跨包） initRouter（私有方法，只能包内） 大小写问题
func InitRouter() {

	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//http://localhost:3000/api/v1/hello
		router.GET("hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": "hello ginblog",
			})

		})

	}
	r.Run(utils.HttpPort)

}
