package model

type Role struct {
	BaseModel
	Code   string `json:"code" gorm:"type:varchar(50);unique;uniqueIndex;not null;comment:角色代码"`
	Name   string `json:"name" gorm:"type:varchar(50);unique;uniqueIndex;not null;comment:角色名称"`
	Enable bool   `json:"enable" gorm:"type:tinyint(4);not null;default:1;comment:启用状态: 1-启用 0-禁用"`
	// 关联信息
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions_permission;"`
}

func (Role) TableName() string {
	return "role"
}
