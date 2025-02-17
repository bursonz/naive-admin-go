package model

type OrderStep struct {
	BaseModel
	// 工单步骤信息
	OrderId uint `json:"orderId" gorm:"type:bigint unsigned;not null;index;comment:工单ID"`
	Sort    int  `json:"sort" gorm:"type:tinyint;not null;comment:步骤顺序"`

	// 执行任务
	// 1: 开锁
	// 2: 关锁
	// 3: 上传图片
	// 4: 状态量
	Task int `json:"task" gorm:"type:tinyint;not null;comment:执行任务类型"`

	// 审核人ID
	ReviewerId *uint `json:"reviewerId" gorm:"type:bigint unsigned;index;comment:审核人ID"`

	// 步骤状态
	// 0: 待执行
	// 1: 已完成
	// 2: 待审核
	// 3: 重新执行
	Status int `json:"status" gorm:"type:tinyint;not null;default:0;index;comment:步骤状态"`

	// 工单步骤内容
	LockId *uint `json:"lockId" gorm:"type:bigint unsigned;index;comment:锁ID"`

	// 锁状态
	// 1: 开锁
	// 2: 关锁
	LockStatus *int `json:"lockStatus" gorm:"type:tinyint;comment:锁状态"`

	ImageUrl *string `json:"imageUrl" gorm:"type:text;comment:图片URL"`
	Comment  *string `json:"comment" gorm:"type:text;comment:步骤内容"`
	SwitchId *string `json:"switchId" gorm:"type:varchar(255);comment:闸刀ID"`
}

func (OrderStep) TableName() string { return "order_step" }
