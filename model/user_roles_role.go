package model

type UserRolesRole struct {
	UserId uint `gorm:"column:userId;type:bigint unsigned;not null;index;comment:用户ID"`
	RoleId uint `gorm:"column:roleId;type:bigint unsigned;not null;index;comment:角色ID"`
}

func (UserRolesRole) TableName() string {
	return "user_roles_role"
}
