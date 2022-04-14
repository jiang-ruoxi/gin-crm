package model

import (
	"gin-crm/model/mysql"
	"gin-crm/utils"
)

//GRole 管理员表
type GRole struct {
	Id         int         `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primary_key"`
	Name       string      `json:"name" gorm:"column:name;type:varchar(255) not null;default:''"`
	Status     uint32      `json:"status" gorm:"column:status;type:tinyint(1) unsigned not null;default:0;index:idx_status"`
	AddTime    utils.XTime `json:"add_time" gorm:"column:add_time;type:timestamp;default:current_timestamp;index:idx_add_time"`
	UpdateTime utils.XTime `json:"update_time" gorm:"column:update_time;type:timestamp;default:current_timestamp on update current_timestamp"`
}

func (GRole) TableName() string {
	return "g_role"
}

//CreateGRole 添加文件数据
func CreateGRole(roleName string, status uint32) bool{
	gRole := GRole{
		Name:   roleName,
		Status: status,
	}
	err:=mysql.DB.Create(&gRole).Error
	if err != nil {
		return false
	}
	return true
}

//修改文件夹名
func UpdateGRole(roleId int, roleName string, status uint32) bool{
	gRole := GRole{
		Name:   roleName,
		Status: status,
	}
	err:=mysql.DB.Model(&GRole{}).Where("id = ?", roleId).Update(gRole).Error
	if err != nil {
		return false
	}
	return true
}

//GetRoleList 获取角色列表
func GetRoleList(pageOffset, pageSize int) (role []GRole, count int, err error) {
	mysql.DB.Limit(pageSize).Offset((pageOffset - 1) * pageSize).Order("id desc").Find(&role)
	var total int
	mysql.DB.Model(&GRole{}).Count(&total)
	return role, total, err
}

//GetRoleInfo 通过fileId获取文件信息
func GetRoleInfo(id int) (role GRole) {
	mysql.DB.First(&role, id)
	return
}

func DeleteDate(id int) bool {
	//删除文件夹信息
	err := mysql.DB.Where("id = ?", id).Delete(GRole{}).Error
	if err != nil {
		return false
	}
	return true
}
