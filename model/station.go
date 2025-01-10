package model

import "time"

type Station struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	AdminUserId int       `json:"adminUserId"`
	Location    string    `json:"location"`
	StationType string    `json:"stationType" comment:"1: 工厂 2: 仓库"`
	CreateTime  time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Station) TableName() string { return "station" }
