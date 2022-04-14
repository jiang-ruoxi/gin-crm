package model

import (
	"fmt"
	"gin-crm/model/mysql"
	"gin-crm/utils"
	_ "github.com/jinzhu/gorm"
)

//GAccess 管理员表
type GAccess struct {
	Id         int         `json:"id" gorm:"column:id;type:int(11) unsigned not null AUTO_INCREMENT;primary_key"`
	ModuleId   int         `json:"module_id" gorm:"column:module_id;type:int(11) unsigned not null;default:0"`
	ModuleName string      `json:"module_name" gorm:"column:module_name;type:varchar(128) not null;default:''"`
	Url        string      `json:"url" gorm:"column:url;type:varchar(128) not null;default:''"`
	Status     int         `json:"status" gorm:"column:status;type:tinyint(1) unsigned not null;default:1;index:idx_status"`
	Type       int         `json:"type" gorm:"column:type;type:tinyint(1) unsigned not null;default:1"`
	Sort       int         `json:"sort" gorm:"column:sort;type:int(11) unsigned not null;default:0"`
	AccessItem []GAccess   `json:"access_item" gorm:"foreignkey:ModuleId;association_foreignkey:Id"`
	AddTime    utils.XTime `json:"add_time" gorm:"column:add_time;type:timestamp;default:current_timestamp;index:idx_add_time"`
	UpdateTime utils.XTime `json:"update_time" gorm:"column:update_time;type:timestamp;default:current_timestamp on update current_timestamp"`
}

func (GAccess) TableName() string {
	return "g_access"
}

//GetAccessList 获取角色列表
func GetAccessList() (access []GAccess) {
	mysql.DB.Preload("AccessItem").Where("module_id=0").Order("sort desc").Find(&access)
	return access
}

//GetAccessTopList 获取角色列表
func GetAccessTopList() (access []GAccess) {
	mysql.DB.Where("module_id=0").Find(&access)
	return access
}

func GetAccessInfo(id int) (access GAccess) {
	mysql.DB.First(&access, id)
	return
}

func GetAccessSonList(id int) (access []GAccess) {
	mysql.DB.Where("module_id=?", id).Find(&access)
	return access
}

func DeleteAccessData(id int) bool {
	//删除文件夹信息
	err := mysql.DB.Where("id = ?", id).Delete(GAccess{}).Error
	if err != nil {
		return false
	}
	return true
}

func CreateGAccess(moduleName, url string, moduleId, status, mType, sort int) bool {
	gAccess := GAccess{
		ModuleName: moduleName,
		Url:        url,
		ModuleId:   moduleId,
		Sort:       sort,
		Status:     status,
		Type:       mType,
	}
	err := mysql.DB.Create(&gAccess).Error
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}
