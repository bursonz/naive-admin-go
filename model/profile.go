package model

type Profile struct {
	ID       int     `json:"id" gorm:"primary_key;auto_increment;type:int(11)"`
	Gender   *int    `json:"gender" gorm:"type:int(11);default:null"`
	Avatar   string  `json:"avatar" gorm:"type:varchar(255);not_null;default:'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80'"`
	Address  *string `json:"address" gorm:"type:varchar(255);default:null"`
	Email    *string `json:"email" gorm:"type:varchar(255);default:null"`
	UserId   int     `json:"userId" gorm:"column:userId;type:int(11);not_null;unique_index"`
	NickName *string `json:"nickName" gorm:"column:nickName;type:varchar(10);default:null"`
	Phone    *string `json:"phone" gorm:"column:phone;type:varchar(13);default:null"`
}

func (Profile) TableName() string {
	return "profile"
}
