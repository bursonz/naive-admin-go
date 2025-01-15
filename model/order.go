package model

type Order struct {
	BaseModel
	// 工单信息
	DispatcherId uint `json:"dispatcherId"` // 派单人id
	OperatorId   uint `json:"operatorId"`   // 操作员id
	StationId    uint `json:"stationId"`    // 站点id
	// 工单状态
	//-2:未通过审批，已驳回
	//-1:已取消
	//0:审批中
	//1:执行中
	//2:审核中
	//3:确认中（审核完成，有序退出挂锁，确认锁状态）
	//4:已完成
	Status int `json:"status"`
	// 工单审批列表
	OrderApprovals []OrderApproval `json:"orderApprovals" gorm:"-"`
	// 工单步骤列表
	OrderSteps []OrderStep `json:"orderSteps" gorm:"-"`
}

func (order Order) TableName() string { return "order" }
