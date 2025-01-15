package model

type OrderStep struct {
	BaseModel
	// 工单步骤信息
	OrderId    uint  `json:"orderId"`    // 工单id
	Sort       int   `json:"sort"`       // 步骤 1,2,3,...10
	Task       int   `json:"task"`       // 执行任务 1:开锁 2:关锁 3:上传图片 4:状态量
	ReviewerId *uint `json:"reviewerId"` // 审核人id  TODO 要不要只保留一个userId，是否需要审核人，还是在工单中显示
	Status     int   `json:"status"`     // 步骤状态 0:待执行 1:已完成 2:待审核 3:重新执行
	// 工单步骤内容
	LockId     *uint   `json:"lockId"`     // 锁id
	LockStatus *int    `json:"lockStatus"` // 锁状态 1:开锁 2:关锁
	ImageUrl   *string `json:"imageUrl"`   // 图片url
	Comment    *string `json:"comment"`    // 状态量
}
