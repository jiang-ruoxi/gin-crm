package model

import (
	"gin-crm/model/mysql"
	"gin-crm/utils"
)

//GRoleAccess 管理员表
type GRoleAccess struct {
	Id         int         `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primary_key"`
	RoleId     int         `json:"role_id" gorm:"column:role_id;type:int(11) unsigned not null;default:0"`
	AccessId   int         `json:"access_id" gorm:"column:access_id;type:int(11) unsigned not null;default:0"`
	AddTime    utils.XTime `json:"add_time" gorm:"column:add_time;type:timestamp;default:current_timestamp;index:idx_add_time"`
	UpdateTime utils.XTime `json:"update_time" gorm:"column:update_time;type:timestamp;default:current_timestamp on update current_timestamp"`
}

func (GRoleAccess) TableName() string {
	return "g_role_access"
}

//CreateGRoleAccess 添加角色权限数据
func CreateGRoleAccess(roleId, accessId int) bool {
	gRoleAccess := GRoleAccess{
		RoleId:   roleId,
		AccessId: accessId,
	}
	err := mysql.DB.Create(&gRoleAccess).Error
	if err != nil {
		return false
	}
	return true
}
