package model

type Role struct {
	BaseModel
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

func (Role) TableName() string {
	return "role"
}
