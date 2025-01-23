package model

type Permission struct {
	BaseModel
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Type        string       `json:"type"`
	ParentId    *uint        `json:"parentId" gorm:"column:parentId"`
	Path        string       `json:"path"`
	Redirect    string       `json:"redirect"`
	Icon        string       `json:"icon"`
	Component   string       `json:"component"`
	Layout      string       `json:"layout"`
	KeepAlive   int          `json:"keepAlive" gorm:"column:keepAlive"`
	Method      string       `json:"method"`
	Description string       `json:"description"`
	Show        int          `json:"show"`
	Enable      int          `json:"enable"`
	Order       int          `json:"order"` // 排序
	Children    []Permission `json:"children" gorm:"-"`
}

func (Permission) TableName() string {
	return "permission"
}
