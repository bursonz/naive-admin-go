package model

type RolePermissionsPermission struct {
	RoleId       uint `gorm:"column:roleId;type:bigint unsigned;not null;index;comment:角色ID"`
	PermissionId uint `gorm:"column:permissionId;type:bigint unsigned;not null;index;comment:权限ID"`
}

func (RolePermissionsPermission) TableName() string {
	return "role_permissions_permission"
}
