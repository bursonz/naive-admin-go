package model

type UserRolesRole struct {
	UserId uint `gorm:"column:userId"`
	RoleId uint `gorm:"column:roleId"`
}

func (UserRolesRole) TableName() string {
	return "user_roles_role"
}
