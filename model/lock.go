package model

import "time"

type Lock struct {
	//gorm.Model
	ID              int       `json:"id"`
	StationId       int       `json:"stationId"`
	AdminId         int       `json:"adminId"`
	SN              int       `json:"sn"`
	Mac             string    `json:"mac"`
	FactoryId       int       `json:"factoryId"`
	CurrentKey      string    `json:"currentKey" comment:"当前密钥"`
	FactoryKey      string    `json:"factoryKey" comment:"出厂密钥"`
	SoftwareVersion string    `json:"softwareVersion"`
	HardwareVersion string    `json:"hardwareVersion"`
	Location        string    `json:"location"`
	Power           int       `json:"power"`
	Description     string    `json:"description"`
	Enable          bool      `json:"enable" comment:"1: 启用 0: 禁用"`
	CreateTime      time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Lock) TableName() string { return "lock" }
