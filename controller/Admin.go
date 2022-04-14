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

func AdminRoleTail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	c.HTML(http.StatusOK, "admin/role_tail.html", gin.H{
		"id": id,
	})
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

func aaa() {
	////1、获取角色id
	//
	//roleId, err := c.GetInt("id")
	//if err != nil {
	//	c.Error("传入参数错误", "/role")
	//	return
	//}
	//
	////2、获取全部的权限
	//
	//access := []models.Access{}
	//models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	//
	////3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	//roleAccess := []models.RoleAccess{}
	//models.DB.Where("role_id=?", roleId).Find(&roleAccess)
	//roleAccessMap := make(map[int]int)
	//for _, v := range roleAccess {
	//	roleAccessMap[v.AccessId] = v.AccessId
	//}
	//
	////4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	//for i := 0; i < len(access); i++ {
	//	if _, ok := roleAccessMap[access[i].Id]; ok {
	//		access[i].Checked = true
	//	}
	//	for j := 0; j < len(access[i].AccessItem); j++ {
	//		if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
	//			access[i].AccessItem[j].Checked = true
	//		}
	//	}
	//}
	////5、渲染权限数据以及角色 Id
	//c.Data["accessList"] = access
	//c.Data["roleId"] = roleId
}
