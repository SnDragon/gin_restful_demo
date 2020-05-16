package api

import (
	"github.com/SnDragon/gin_restful_demo/dao"
	"github.com/SnDragon/gin_restful_demo/model"
	"github.com/SnDragon/gin_restful_demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建用户
func CreateUser(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ERROR,
			"msg":  err.Error(),
		})
		return
	}
	dao.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "success",
		"data": user,
	})
}

// 更新用户
func UpdateUser(c *gin.Context) {
	var status = http.StatusOK
	var code int
	var msg string
	var data *model.User
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		status, code, msg = http.StatusBadRequest, ERROR, err.Error()
	} else {
		id,err := utils.GetInt(c.Param("id"))
		if err == nil {
			data = dao.UpdateUserById(id, &user)
		}
		if err != nil || data == nil || data.ID == 0 {
			code, msg, data = ERROR, "error", nil
		}
	}

	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var code = SUCCESS
	var msg = "success"
	userId := c.Param("id")
	deleted := dao.DeleteUserById(userId)
	if !deleted {
		code, msg = ERROR, "error"
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

// 获取用户
func GetUser(c *gin.Context) {
	var code = SUCCESS
	var msg = "success"
	var data *model.User
	id,err := utils.GetInt(c.Param("id"))
	if err == nil {
		data = dao.GetUserById(id)
	}
	if err != nil || data == nil || data.ID == 0 {
		code, msg, data = ERROR, "error", nil
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	page, err1 := utils.GetInt(c.DefaultQuery("page", "1"))
	limit, err2 := utils.GetInt(c.DefaultQuery("limit", "10"))

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERROR, "msg": "param error"})
		return
	}
	total, users := dao.GetUsers(page, limit)
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "success",
		"data": gin.H{
			"total": total,
			"list":  users,
		},
	})
}
