package controller

import (
	"gin-crm/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetManagerList(c *gin.Context) {
	//获取用户信息
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	roleList, total, _ := model.GetManagerList(page, limit)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": total,
		"data":  roleList,
	})
}

func DoManagerAddData(c *gin.Context) {
	nickname := c.PostForm("nick_name")
	password := c.PostForm("password")
	mobile := c.PostForm("mobile")
	email := c.PostForm("email")
	sex, _ := strconv.Atoi(c.PostForm("sex"))
	sexUint32 := uint32(sex)
	roleId, _ := strconv.Atoi(c.PostForm("role_id"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	statusUint32 := uint32(status)

	managerInfo := model.GetManagerInfo(nickname)
	if managerInfo.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "该用户名已经存在",
		})
		return
	}

	res := model.CreateGManager(roleId, nickname, password, mobile, email, statusUint32, sexUint32)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "创建成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "创建失败",
		})
	}

}



