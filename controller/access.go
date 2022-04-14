package controller

import (
	"gin-crm/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAccessInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("access_id"))
	info := model.GetAccessInfo(id)


	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": id,
		"data":    info,
	})

}

func DoAccessAddData(c *gin.Context) {
	url := c.PostForm("url")
	moduleName := c.PostForm("module_name")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	mType, _ := strconv.Atoi(c.PostForm("type"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	moduleId, _ := strconv.Atoi(c.PostForm("module_id"))

	res := model.CreateGAccess(moduleName, url, moduleId, status, mType, sort)
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

func DoDeleteAccessData(c *gin.Context) {
	//获取用户信息
	id, _ := strconv.Atoi(c.PostForm("id"))

	accessList := model.GetAccessSonList(id)
	if len(accessList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "存在子模块，禁止删除!",
			"data":    "",
		})
		return
	}

	if ok := model.DeleteAccessData(id); ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "操作成功",
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    10001,
		"message": "操作失败",
		"data":    "",
	})
}

func DoSaveAuth(c *gin.Context)  {
	roleId, _ := strconv.Atoi(c.PostForm("role_id"))
	info := model.GetRoleInfo(roleId)
	if info.Id !=  roleId{
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "非法请求！",
		})
		return
	}

	//accessNode:= c.PostFormMap("access_node")
	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": c.Request.ParseForm(),
	})
	//for _, v := range accessNode {
	//	accessId, _ := strconv.Atoi(v)
	//	model.CreateGRoleAccess(roleId,accessId)
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    10000,
	//	"message": "操作成功！",
	//})
}
