package model

type OrderApproval struct {
	BaseModel
	// 工单ID
	OrderId uint `json:"orderId" gorm:"type:bigint unsigned;not null;index;comment:工单ID"`

	// 审批人ID
	ApproverId uint `json:"approverId" gorm:"type:bigint unsigned;not null;index;comment:审批人ID"`

	// 审批状态
	// -1: 拒绝
	//  0: 待审批
	//  1: 已审批
	Status int `json:"status" gorm:"type:tinyint;not null;default:0;index;comment:审批状态"`

	// 审批意见
	Comment *string `json:"comment" gorm:"type:text;comment:审批意见"`

	// 审批顺序
	Sort int `json:"sort" gorm:"type:tinyint;not null;comment:审批顺序"`
}

func (OrderApproval) TableName() string {
	return "order_approval"
}
