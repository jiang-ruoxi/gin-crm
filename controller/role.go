package controller

import (
	"gin-crm/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRoleList(c *gin.Context) {
	//获取用户信息
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	roleList, total, _ := model.GetRoleList(page, limit)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": total,
		"data":  roleList,
	})
}

func DoDeleteData(c *gin.Context) {
	//获取用户信息
	id, _ := strconv.Atoi(c.PostForm("id"))

	info := model.GetRoleInfo(id)
	if info.Id == id {
		if ok := model.DeleteDate(id); ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
				"data":    "",
			})
		}
	}
	if info.Id != id {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "failed",
			"data":    "",
		})
	}
}

func DoAddData(c *gin.Context) {
	role := c.PostForm("role_name")
	status, _ := strconv.Atoi(c.PostForm("status"))
	statusUint32 := uint32(status)
	res := model.CreateGRole(role,statusUint32)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "创建成功",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "创建失败",
		})
	}

}

func DoEditData(c *gin.Context) {
	role := c.PostForm("role_name")
	roleId, _ := strconv.Atoi(c.PostForm("role_id"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	statusUint32 := uint32(status)
	roleId = int(roleId)
	res := model.UpdateGRole(roleId,role,statusUint32)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "success",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "failed",
		})
	}

}


func GetRoleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("role_id"))
	info := model.GetRoleInfo(id)


		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": id,
			"data":    info,
		})

}
