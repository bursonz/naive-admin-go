package model

type Order struct {
	BaseModel
	// 工单信息
	DispatcherId uint `json:"dispatcherId" gorm:"type:bigint unsigned;not null;index;comment:派单人ID"`
	OperatorId   uint `json:"operatorId" gorm:"type:bigint unsigned;not null;index;comment:操作员ID"`
	StationId    uint `json:"stationId" gorm:"type:bigint unsigned;not null;index;comment:站点ID"`

	// 工单状态
	// -2: 未通过审批，已驳回
	// -1: 已取消
	//  0: 审批中
	//  1: 执行中
	//  2: 审核中
	//  3: 确认中（审核完成，有序退出挂锁，确认锁状态）
	//  4: 已完成
	Status int `json:"status" gorm:"type:tinyint(4);not null;default:0;index;comment:工单状态"`

	// 工单审批列表
	OrderApprovals []OrderApproval `json:"orderApprovals" gorm:"foreignKey:OrderId"`

	// 工单步骤列表
	OrderSteps []OrderStep `json:"orderSteps" gorm:"foreignKey:OrderId"`

	// 工单内容
	Content string `json:"content" gorm:"type:text;comment:工单内容"`
}

func (order Order) TableName() string { return "order" }
