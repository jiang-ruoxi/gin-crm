package model

import (
	"fmt"
	"gin-crm/model/mysql"
	"gin-crm/utils"
	_ "github.com/jinzhu/gorm"
)

//GManager 管理员表
type GManager struct {
	Id         int         `json:"id"`
	RoleId     int         `json:"role_id"`
	NickName   string      `json:"nick_name"`
	Password   string      `json:"password"`
	Mobile     string      `json:"mobile"`
	Email      string      `json:"email"`
	Status     int         `json:"status"`
	Sex        uint32      `json:"sex"`
	IsSuper    uint32      `json:"is_super"`
	Role       GRole       `gorm:"foreignkey:Id;association_foreignkey:RoleId"`
	AddTime    utils.XTime `json:"add_time" gorm:"column:add_time;type:timestamp;default:current_timestamp;index:idx_add_time"`
	UpdateTime utils.XTime `json:"update_time" gorm:"column:update_time;type:timestamp;default:current_timestamp on update current_timestamp"`
	RoleName   string      `json:"role_name" gorm:"-"` // 忽略本字段
}

func (GManager) TableName() string {
	return "g_manager"
}

//GetManagerList 获取角色列表
func GetManagerList(pageOffset, pageSize int) (role []GManager, count int, err error) {
	mysql.DB.Preload("Role").Limit(pageSize).Offset((pageOffset - 1) * pageSize).Order("id desc").Find(&role)
	var total int
	mysql.DB.Model(&GManager{}).Count(&total)
	return role, total, err
}

//GetManagerInfo
func GetManagerInfo(nickName string) (manager GManager) {
	mysql.DB.Where("nick_name = ?", nickName).First(&manager)
	return
}

//CreateGManager 添加管理员
func CreateGManager(roleId int, nickName, password, mobile, email string, status int, sex uint32) bool {
	gManager := GManager{
		RoleId:   roleId,
		NickName: nickName,
		Password: utils.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   status,
		Sex:      sex,
	}
	fmt.Println(gManager)
	err := mysql.DB.Create(&gManager).Error
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}

func UpdateGManagerStatus(managerId, status int) bool {
	gManager := GManager{
		Status: status,
	}
	err := mysql.DB.Model(&GManager{}).Where("id = ?", managerId).Update(gManager).Error
	if err != nil {
		return false
	}
	return true
}

func GetGManagerInfo(managerId int) (manager GManager) {
	mysql.DB.Where("id=?", managerId).Find(&manager)
	return manager
}

func IsSuperManager(managerId int) bool {
	var manager GManager
	mysql.DB.Where("id=? and is_super = 1", managerId).Find(&manager)
	if manager.Id != managerId {
		return false
	}
	return true
}

func IsExistNickName(nickName string) bool {
	var manager GManager
	mysql.DB.Where("nick_name = ?", nickName).First(&manager)
	if manager.NickName == nickName {
		return true
	}
	return false
}

func DeleteManagerData(id int) bool {
	err := mysql.DB.Where("id = ?", id).Delete(GManager{}).Error
	if err != nil {
		return false
	}
	return true
}
