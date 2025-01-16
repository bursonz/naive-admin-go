package model

import "time"

type Log struct {
	ID         int       `json:"id"`
	UserId     int       `json:"userId"`
	TargetItem string    `json:"targetItem"`
	TargetId   int       `json:"targetId"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
}
