package controller

import (
	"gin-crm/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetManagerList 获取管理员列表
func GetManagerList(c *gin.Context) {
	//获取分页参数
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	//查询数据库
	managerList, total, _ := model.GetManagerList(page, limit)

	//处理数据组合展示
	roleAccess := []model.GManager{}
	for _, item := range managerList {
		arr := model.GManager{
			Id:       item.Id,
			NickName: item.NickName,
			Status:   item.Status,
			IsSuper:  item.IsSuper,
			AddTime:  item.AddTime,
			RoleName: item.Role.Name,
		}
		roleAccess = append(roleAccess, arr)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": total,
		"data":  roleAccess,
	})
}

//GetManagerInfo 获取管理员详情
func GetManagerInfo(c *gin.Context) {
	managerId, _ := strconv.Atoi(c.PostForm("manager_id"))
	managerInfo := model.GetGManagerInfo(managerId)
	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": managerId,
		"data":    managerInfo,
	})
}

//DoManagerAddData 执行添加管理员操作
func DoManagerAddData(c *gin.Context) {
	//获取form表单参数
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")
	nickname := c.PostForm("nick_name")
	password := c.PostForm("password")
	sex, _ := strconv.Atoi(c.PostForm("sex"))
	sexUint32 := uint32(sex)
	status, _ := strconv.Atoi(c.PostForm("status"))
	roleId, _ := strconv.Atoi(c.PostForm("role_id"))

	//判断是否存在了用户名
	if isBool := model.IsExistNickName(nickname); isBool {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "该用户名已经存在",
		})
		return
	}

	if isSuccess := model.CreateGManager(roleId, nickname, password, mobile, email, status, sexUint32); isSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "创建成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    10001,
		"message": "创建失败",
	})
}

//DoEditManagerStatus 执行修改状态的操作
func DoEditManagerStatus(c *gin.Context) {
	managerId, _ := strconv.Atoi(c.PostForm("manager_id"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if isBool := model.UpdateGManagerStatus(managerId, status); isBool {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "操作成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    10001,
		"message": "操作失败",
	})
}

//DoManagerDeleteData 执行删除数据操作
func DoManagerDeleteData(c *gin.Context) {
	managerId, _ := strconv.Atoi(c.PostForm("manager_id"))
	isBool := model.IsSuperManager(managerId)
	if isBool {
		c.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": "超级管理员禁止删除",
		})
		return
	}
	if ok := model.DeleteManagerData(managerId); ok {
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
