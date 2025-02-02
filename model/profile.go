package model

type Profile struct {
	BaseModel
	Gender   *int    `json:"gender" gorm:"type:int(11);default:null;comment:性别: 0-未知 1-男 2-女"`
	Avatar   string  `json:"avatar" gorm:"type:varchar(255);not null;default:'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80';comment:头像URL"`
	Address  *string `json:"address" gorm:"type:varchar(255);default:null;comment:详细地址"`
	Email    *string `json:"email" gorm:"type:varchar(255);default:null;comment:电子邮箱"`
	UserId   uint    `json:"userId" gorm:"column:userId;type:bigint unsigned;not null;uniqueIndex;comment:关联用户ID"`
	NickName *string `json:"nickName" gorm:"column:nickName;type:varchar(32);default:null;comment:用户昵称"`
	Phone    *string `json:"phone" gorm:"column:phone;type:varchar(20);default:null;comment:手机号码"`
}

func (Profile) TableName() string {
	return "profile"
}
