package model

type Station struct {
	BaseModel
	Code        string `json:"code"`
	Name        string `json:"name"`
	AdminUserId uint   `json:"adminUserId"`
	Location    string `json:"location"`
	StationType string `json:"stationType" comment:"1: 工厂 2: 仓库"`
}

func (Station) TableName() string { return "station" }
