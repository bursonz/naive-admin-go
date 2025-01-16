package api

type LockStatus int

const (
	LockStatusInit     LockStatus = 0 // 初始值，不使用
	LockStatusLocked   LockStatus = 1 // 开
	LockStatusUnlocked LockStatus = 2 // 关
)

type OrderStatus int

// 工单    0待审批，1审批中，2待执行，3执行中，4待审核，5待确认，
// 审批状态 0待审批，1审批中，2审批通过，3审批驳回
// 执行状态 0待执行，4执行中，6审核中，7审核驳回重新执行中，8确认中，9已完成
const (
	OrderStatusInit        OrderStatus = 0 // 初始值，不使用
	OrderStatusApproving   OrderStatus = 1 // 审批中，待审批：工单创建完成后的初始状态
	OrderStatusApproved    OrderStatus = 2 // 已审批，审批通过
	OrderStatusRejected    OrderStatus = 3 // 已审批，审批驳回，已取消
	OrderStatusExecuting   OrderStatus = 4 // 执行中，待执行
	OrderStatusExecuted    OrderStatus = 5 // 已执行
	OrderStatusReviewing   OrderStatus = 6 // 审核中
	OrderStatusReExecuting OrderStatus = 7 // 审核驳回，重新执行
	OrderStatusFinished    OrderStatus = 8 // 审核通过，确认中：需要倒序遍历工单步骤，并依次确认关锁
	OrderStatusCompleted   OrderStatus = 9 // 已完成
)

type OrderApprovalStatus int

// 0待审批，1审批中，2审批通过，3审批驳回
const (
	OrderApprovalStatusInit      OrderApprovalStatus = 0 // 初始值，不使用
	OrderApprovalStatusApproving OrderApprovalStatus = 1 // 审批中，待审批
	OrderApprovalStatusApproved  OrderApprovalStatus = 2 // 已审批，审批通过
	OrderApprovalStatusRejected  OrderApprovalStatus = 3 // 已审批，审批驳回
)

type OrderStepStatus int

// 执行中，已执行，已确认，待审核，审核中，已审核，审核驳回待重新执行
const (
	OrderStepStatusInit        OrderStepStatus = 0 // 等待审批
	OrderStepStatusExecuting   OrderStepStatus = 4 // 执行中
	OrderStepStatusExecuted    OrderStepStatus = 5 // 已执行
	OrderStepStatusReviewing   OrderStepStatus = 6 // 审核中
	OrderStepStatusReExecuting OrderStepStatus = 7 // 审核驳回，重新执行
	OrderStepStatusFinished    OrderStepStatus = 8 // 审核通过
	OrderStepStatusCompleted   OrderStepStatus = 9 // 已完成
)

type OrderStepTask int

const (
	OrderStepTaskInit                OrderStepTask = 0 // 初始值，不使用
	OrderStepTaskLock                OrderStepTask = 1 // 关锁
	OrderStepTaskUnlock              OrderStepTask = 2 // 开锁
	OrderStepTaskUploadConfirmImage  OrderStepTask = 3 // 上传确认图片
	OrderStepTaskRetrieveStatusValue OrderStepTask = 4 // 获取状态量
)
