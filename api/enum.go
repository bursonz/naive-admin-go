package api

type OrderStatus int

// 工单状态
const (
	OrderRejected   = -1
	OrderApproving  = 1
	OrderExecuting  = 3
	OrderReviewing  = 6
	OrderConfirming = 8
	OrderFinished   = 10
)

// 工单    0待审批，1审批中，2待执行，3执行中，4待审核，5待确认，
// 审批状态 0待审批，1审批中，2审批通过，3审批驳回
// 执行状态 0待执行，4执行中，6审核中，7审核驳回重新执行中，8确认中，9已完成
const (
	OrderStatusRejected   = -1 // 已驳回，已审批，审批驳回，已取消
	OrderStatusInit       = 0  // 初始值，不使用
	OrderStatusApproving  = 1  // 审批中，待审批：工单创建完成后的初始状态
	OrderStatusApproved   = 2  // 已审批，审批通过
	OrderStatusExecuting  = 3  // 执行中，待执行
	OrderStatusReExecute  = 4  // 审核驳回，重新执行
	OrderStatusExecuted   = 5  // 已执行
	OrderStatusReviewing  = 6  // 审核中
	OrderStatusReviewed   = 7  // 审核通过，
	OrderStatusConfirming = 8  // 确认中：需要倒序遍历工单步骤，并依次确认关锁
	OrderStatusConfirmed  = 9  // 已确认
	OrderStatusFinished   = 10 // 已完成
)

type OrderApprovalStatus int

// 工单审批状态
// 0待审批，1审批中，2审批通过，3审批驳回
const (
	OrderApprovalStatusRejected  = -1 // 已审批，审批驳回
	OrderApprovalStatusApproving = 1  // 审批中，待审批
	OrderApprovalStatusApproved  = 2  // 已审批，审批通过
)

type OrderStepStatus int

// 工单步骤状态
// 执行中，已执行，已确认，待审核，审核中，已审核，审核驳回待重新执行
const (
	OrderStepStatusExecuting  = 3  // 执行中
	OrderStepStatusReExecute  = 4  // 重新执行，审核驳回
	OrderStepStatusExecuted   = 5  // 已执行
	OrderStepStatusReviewing  = 6  // 审核中
	OrderStepStatusReviewed   = 7  // 已审核，审核通过
	OrderStepStatusConfirming = 8  // 确认中：需要倒序遍历工单步骤，并依次确认关锁
	OrderStepStatusConfirmed  = 9  // 已确认，标记为已完成（可以在前端使用额外字段暂存）
	OrderStepStatusFinished   = 10 // 已完成
)

type OrderStepTask int

const (
	OrderStepTaskInit                = 0 // 初始值，不使用
	OrderStepTaskLock                = 1 // 关锁
	OrderStepTaskUnlock              = 2 // 开锁
	OrderStepTaskUploadConfirmImage  = 3 // 上传确认图片
	OrderStepTaskRetrieveStatusValue = 4 // 获取状态量
)

// 锁状态
type LockStatus int

const (
	LockStatusLocked = 0 // 锁关
	LockStatusUnLock = 1 // 锁开
)
