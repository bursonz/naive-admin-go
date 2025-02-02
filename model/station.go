package model

type Station struct {
	BaseModel
	Code        string `json:"code" gorm:"type:varchar(64);index;not null;comment:唯一标识"`
	Name        string `json:"name" gorm:"type:varchar(64);index;not null;comment:名称"`
	AdminUserId uint   `json:"adminUserId" gorm:"type:bigint unsigned;not null;index;comment:管理员ID"`
	Location    string `json:"location" gorm:"type:varchar(255);default:null;comment:位置"`
	StationType string `json:"stationType" gorm:"type:varchar(64);default:null;comment:站点类型"`
}

func (Station) TableName() string { return "station" }
