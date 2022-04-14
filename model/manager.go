package model

import (
	"fmt"
	"gin-crm/model/mysql"
	"gin-crm/utils"
	_ "github.com/jinzhu/gorm"
)

//GManager 管理员表
type GManager struct {
	Id         int         `json:"id" gorm:"column:id;type:int(11) unsigned not null AUTO_INCREMENT;primary_key"`
	RoleId     int         `json:"role_id" gorm:"column:role_id;type:int(11) unsigned not null;default:0"`
	NickName   string      `json:"nick_name" gorm:"column:nick_name;type:varchar(255) not null;default:''"`
	Password   string      `json:"password" gorm:"column:password;type:char(32) not null;default:''"`
	Mobile     string      `json:"mobile" gorm:"column:mobile;type:char(11) not null;default:''"`
	Email      string      `json:"email" gorm:"column:email;type:varchar(255) not null;default:''"`
	Status     uint32      `json:"status" gorm:"column:status;type:tinyint(1) unsigned not null;default:1;index:idx_status"`
	Sex        uint32      `json:"sex" gorm:"column:sex;type:tinyint(1) unsigned not null;default:1"`
	IsSuper    uint32      `json:"is_super" gorm:"column:is_supper;type:tinyint(1) unsigned not null;default:0"`
	Role       GRole       `gorm:"foreignkey:Id;association_foreignkey:RoleId"`
	AddTime    utils.XTime `json:"add_time" gorm:"column:add_time;type:timestamp;default:current_timestamp;index:idx_add_time"`
	UpdateTime utils.XTime `json:"update_time" gorm:"column:update_time;type:timestamp;default:current_timestamp on update current_timestamp"`
}

func (GManager) TableName() string {
	return "g_manager"
}

//GetManagerList 获取角色列表
func GetManagerList(pageOffset, pageSize int) (role []GManager, count int, err error) {
	mysql.DB.Limit(pageSize).Offset((pageOffset - 1) * pageSize).Order("id desc").Find(&role)
	var total int
	mysql.DB.Model(&GManager{}).Count(&total)
	return role, total, err
}

//GetManagerInfo 通过fileId获取文件信息
func GetManagerInfo(nickName string) (manager GManager) {
	mysql.DB.Where("nick_name = ?", nickName).First(&manager)
	return
}

//CreateGManager 添加管理员
func CreateGManager(roleId int, nickName, password, mobile, email string, status, sex uint32) bool {
	gManager := GManager{
		RoleId:   roleId,
		NickName: nickName,
		Password: utils.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   status,
		Sex:      sex,
	}
	err := mysql.DB.Create(&gManager).Error
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}
