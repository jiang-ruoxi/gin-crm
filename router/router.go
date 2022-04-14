package router

import (
	"gin-crm/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//后台路由组
	admin := router.Group("admin")
	admin.Use()
	{
		admin.GET("/index", controller.AdminIndex)
		admin.GET("/home", controller.AdminHome)
		admin.GET("/role", controller.AdminRole)
		admin.GET("/auth", controller.AdminRoleAuth)

		admin.GET("/access", controller.AdminAccess)
		admin.POST("/access/add", controller.DoAccessAddData)
		admin.POST("/access/info", controller.GetAccessInfo)
		admin.POST("/access/delete", controller.DoDeleteAccessData)
		admin.POST("/auth/do", controller.DoSaveAuth)

		admin.GET("/role/list", controller.GetRoleList)
		admin.POST("/role/add", controller.DoAddData)
		admin.POST("/role/delete", controller.DoDeleteData)
		admin.POST("/role/info", controller.GetRoleInfo)
		admin.POST("/role/edit", controller.DoEditData)

		admin.GET("/manager", controller.AdminManager)
		admin.GET("/manager/list", controller.GetManagerList)
		admin.POST("/manager/info", controller.GetManagerInfo)
		admin.POST("/manager/add", controller.DoManagerAddData)
		admin.POST("/manager/edit/status", controller.DoEditManagerStatus)
		admin.POST("/manager/delete", controller.DoManagerDeleteData)

	}

	router.LoadHTMLGlob("view/**/*")
	router.Static("/static", "./static")
	return router
}
