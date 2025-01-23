package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Enable   bool   `json:"enable"`
}

func (User) TableName() string {
	return "user"
}
