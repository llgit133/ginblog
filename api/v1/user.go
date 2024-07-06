package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserExist(c *gin.Context) {

}

func AddUser(c *gin.Context) {

	var data model.User
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	// 返回json 数据 code ,data,msg
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetUsers(c *gin.Context) {

	// 从前端获取参数 pageSize pageSum
	// String 转为int
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageSum, _ := strconv.Atoi(c.Query("pageNum"))

	// gorm 框架 当为-1时，查询所有
	if pageSize == 0 {
		pageSize = -1
	}

	if pageSum == 0 {
		pageSum = -1
	}
	users := model.GetUsers(pageSize, pageSum)
	code := errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})

}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)

	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		// 直接返回 不执行
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
