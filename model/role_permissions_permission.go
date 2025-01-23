package model

type RolePermissionsPermission struct {
	RoleId       uint `gorm:"column:roleId"`
	PermissionId uint `gorm:"column:permissionId"`
}

func (RolePermissionsPermission) TableName() string {
	return "role_permissions_permission"
}
