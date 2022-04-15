package controller

import (
	"gin-crm/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminHome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/home.html", gin.H{})
}

func AdminIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}

func AdminRole(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role.html", gin.H{})
}

func AdminManager(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager.html", gin.H{})
}

func AdminRoleAuth(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("role_id"))
	info := model.GetRoleInfo(roleId)
	if info.Id != roleId {
		c.Redirect(http.StatusMovedPermanently,"/admin/role")
		return
	}
	accessList := model.GetAccessList()
	c.HTML(http.StatusOK, "admin/auth.html", gin.H{
		"role_id":     roleId,
		"access_list": accessList,
	})
}

func AdminAccess(c *gin.Context) {
	accessList := model.GetAccessList()
	accessTopList := model.GetAccessTopList()
	c.HTML(http.StatusOK, "admin/access.html", gin.H{
		"access_list":     accessList,
		"access_top_list": accessTopList,
	})
}
