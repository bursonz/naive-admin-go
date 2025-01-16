package model

type OrderApproval struct {
	BaseModel
	//工单审批信息
	OrderId    uint    `json:"orderId"`    // 工单id
	ApproverId uint    `json:"approverId"` // 审批人id
	Status     int     `json:"status"`     // 审批状态 -1:拒绝 0:待审批 1:已审批
	Comment    *string `json:"comment"`    // 审批意见
	Sort       int     `json:"sort"`       // 审批顺序 1,2,3,...10
}

func (OrderApproval) TableName() string {
	return "order_approval"
}
