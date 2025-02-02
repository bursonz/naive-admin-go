package model

type Permission struct {
	BaseModel
	Name        string       `json:"name" gorm:"type:varchar(255);not null;comment:权限名称"`
	Code        string       `json:"code" gorm:"type:varchar(50);unique;uniqueIndex;not null;comment:权限代码"`
	Type        string       `json:"type" gorm:"type:varchar(255);not null;comment:权限类型: MENU-菜单 BUTTON-按钮"`
	ParentId    *uint        `json:"parentId" gorm:"column:parentId;type:bigint unsigned;default:null;comment:父级ID"`
	Path        string       `json:"path" gorm:"type:varchar(255);default:null;comment:路由路径"`
	Redirect    string       `json:"redirect" gorm:"type:varchar(255);default:null;comment:重定向路径"`
	Icon        string       `json:"icon" gorm:"type:varchar(255);default:null;comment:图标"`
	Component   string       `json:"component" gorm:"type:varchar(255);default:null;comment:组件路径"`
	Layout      string       `json:"layout" gorm:"type:varchar(255);default:null;comment:布局"`
	KeepAlive   int          `json:"keepAlive" gorm:"column:keepAlive;type:tinyint(4);default:null;comment:是否缓存: 1-是 0-否"`
	Method      string       `json:"method" gorm:"type:varchar(255);default:null;comment:请求方法"`
	Description string       `json:"description" gorm:"type:varchar(255);default:null;comment:描述"`
	Show        int          `json:"show" gorm:"type:tinyint(4);not null;default:1;comment:是否展示在页面菜单: 1-是 0-否"`
	Enable      int          `json:"enable" gorm:"type:tinyint(4);not null;default:1;comment:启用状态: 1-启用 0-禁用"`
	Order       int          `json:"order" gorm:"type:int(11);default:null;comment:排序"`
	Children    []Permission `json:"children" gorm:"-"`
}

func (Permission) TableName() string {
	return "permission"
}
