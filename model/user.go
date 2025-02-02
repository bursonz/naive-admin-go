package model

type User struct {
	BaseModel
	Username string `json:"username" gorm:"type:varchar(50);unique;uniqueIndex;not null;comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);not null;comment:密码"`
	Enable   bool   `json:"enable" gorm:"type:tinyint(4);not null;default:1;comment:启用状态: 1-启用 0-禁用"`
	// 关联信息
	Roles []Role `json:"roles" gorm:"many2many:user_roles_role;"`
}

func (User) TableName() string {
	return "user"
}
