package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

type orderStep struct{}

var OrderStep = &orderStep{}

func (orderStep) List(c *gin.Context) {
	var data inout.OrderStepListRes
	// 查询条件 TODO 其他条件
	var orderId = c.DefaultQuery("orderId", "")
	var task = c.DefaultQuery("task", "")
	var operatorId = c.DefaultQuery("operatorId", "")
	var reviewerId = c.DefaultQuery("reviewerId", "")
	var lockId = c.DefaultQuery("lockId", "")
	var status = c.DefaultQuery("status", "")
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var deleted = c.DefaultQuery("deleted", "")
	// 条件查询
	var orm = db.Dao
	if deleted != "" {
		orm = orm.Unscoped()
	}
	orm = orm.Model(&model.OrderStep{})
	if orderId != "" {
		orm = orm.Where("order_id like ?", "%"+orderId+"%")
	}
	if task != "" {
		orm = orm.Where("task like ?", "%"+task+"%")
	}
	if operatorId != "" {
		orm = orm.Where("operator_id like ?", "%"+operatorId+"%")
	}
	if reviewerId != "" {
		orm = orm.Where("reviewer_id like ?", "%"+reviewerId+"%")
	}
	if status != "" {
		orm = orm.Where("status = ?", status)
	}
	if lockId != "" {
		orm = orm.Where("lock_id like ?", "%"+lockId+"%")
	}
	// 查询总数
	orm.Count(&data.Total)
	// 分页查询
	if pageNo < 1 { // pageNo 小于1 时，查询所有
		orm.Find(&data.PageData)
	} else {
		orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("sort asc").Find(&data.PageData)
	}
	Resp.Succ(c, data)
}

func (orderStep) Add(c *gin.Context) {
	var params inout.AddOrderStepReq

	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		var newOrderStep = model.OrderStep{
			OrderId:    params.OrderId,
			Sort:       params.Sort,
			Task:       params.Task,
			ReviewerId: params.ReviewerId,
			Status:     params.Status,
			LockId:     params.LockId,
			LockStatus: params.LockStatus,
			ImageUrl:   params.ImageUrl,
			Comment:    params.Comment,
		}
		if err := tx.Model(&model.OrderStep{}).Create(&newOrderStep).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, err)
	}
}

func (orderStep) Update(c *gin.Context) {
	var params inout.PatchOrderStepReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		var err error
		orm := tx.Model(&model.OrderStep{}).Where("id = ?", params.Id)
		if params.Task != nil {
			err = orm.Update("task", params.Task).Error
			if err != nil {
				return err
			}
		}
		if params.Sort != nil {
			err = orm.Update("sort", params.Sort).Error
			if err != nil {
				return err
			}
		}
		if params.ImageUrl != nil {
			err = orm.Update("image_url", params.ImageUrl).Error
			if err != nil {
				return err
			}
		}
		if params.Comment != nil {
			err = orm.Update("comment", params.Comment).Error
			if err != nil {
				return err
			}
		}
		if params.LockId != nil {
			err = orm.Update("lock_id", params.LockId).Error
			if err != nil {
				return err
			}
		}
		if params.ReviewerId != nil {
			err = orm.Update("reviewer_id", params.ReviewerId).Error
			if err != nil {
				return err
			}
		}
		if params.LockStatus != nil {
			err = orm.Update("lock_status", params.LockStatus).Error
			if err != nil {
				return err
			}
			// TODO 应该删除这部分，改为前端实现工单结束
			// 如果锁状态都为关，则更新工单步骤状态为已完成
			//if *params.LockStatus == LockStatusLocked {
			//	var currentOrder model.Order
			//	tx.Model(&model.Order{}).
			//		Where("id =?", params.OrderId).
			//		Select("status").
			//		Find(&currentOrder)
			//	if currentOrder.Status == OrderConfirming {
			//		// 查询所有未关锁的步骤数量
			//		var unlockCount int64
			//		tx.Model(&model.OrderStep{}).
			//			Where("order_id = ?", params.OrderId).
			//			Where("task =?", OrderStepTaskUnlock).
			//			Where("lock_status =?", LockStatusUnLock).
			//			Count(&unlockCount)
			//		if unlockCount == 0 { // 没有，即锁全部关闭
			//			tx.Model(&model.Order{}).
			//				Where("order_id =?", params.OrderId).
			//				Update("status", OrderFinished)
			//		}
			//	}
			//}
		}
		if params.Status != nil {
			// 判断工单状态，如果工单状态为确认中，则无需响应，直接返回
			var currentOrder model.Order
			tx.Model(&model.Order{}).Where("id =?", params.OrderId).Select("status").Find(&currentOrder)
			switch currentOrder.Status {
			case OrderRejected:
				return errors.New("工单已被驳回，不能执行！")
			case OrderApproving:
				return errors.New("工单正在审批中，不能执行！")
			case OrderExecuting:
				switch *params.Status {
				case OrderStepStatusExecuted:
					// 执行中 -> 已执行:
					// 判断工单步骤类型是否为审核图片上传或审核状态量
					switch *params.Task {
					case OrderStepTaskUploadConfirmImage, OrderStepTaskRetrieveStatusValue:
						// 如果是，则更新工单步骤状态为待审核, 并更新工单状态为审核中
						orm.Update("status", OrderStepStatusReviewing)
						tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderReviewing)
					case OrderStepTaskUnlock, OrderStepTaskLock:
						// 如果不是，则更新工单步骤状态为执行中
						orm.Update("status", OrderStepStatusExecuted)
					}
				default:
				}
			case OrderReviewing:
				switch *params.Status {
				case OrderStepStatusReExecute: // 审核中 -> 重新执行
					// 更新工单步骤状态为重新执行
					orm.Update("status", OrderStepStatusReExecute)
					// 更新工单步骤状态为执行中
					tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderExecuting)

					// TODO删除工单步骤中的记录
					// 判断类型是否为审核图片上传
					//if *params.Task == OrderStepTaskUploadConfirmImage {
					//	// 删除工单步骤中的文件及记录
					//	var filename string
					//	tx.Model(&model.OrderStep{}).
					//		Where("id =?", params.Id).Find(&filename)
					//}
				case OrderStepStatusReviewed: // 审核中 -> 已审核
					orm.Update("status", OrderStepStatusReviewed)
					var totalCount, currentCount int64
					tx.Model(&model.OrderStep{}).
						Where("order_id =?", params.OrderId).
						Count(&totalCount).
						Where("status in (?)", []int{OrderStepStatusReviewed, OrderStepStatusExecuted}).
						Count(&currentCount)
					// 如果所有步骤都已审核或执行，则更新工单状态为确认中
					if totalCount == currentCount {
						tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderConfirming)
					} else { // 否则更新工单状态为执行中
						tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderExecuting)
					}
					// 判断工单步骤类型
					//switch *params.Task {
					//case OrderStepTaskUploadConfirmImage: // 审核图片上传
					//	var count int64
					//	tx.Model(&model.OrderStep{}).
					//		Where("order_id =?", params.OrderId).
					//		Where("task =?", OrderStepTaskRetrieveStatusValue).
					//		Count(&count)
					//	if count != 0 {
					//		//tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderExecuting)
					//		break
					//	}
					//	fallthrough
					//case OrderStepTaskRetrieveStatusValue: // 审核状态量
					//	//tx.Model(&model.Order{}).Where("id =?", params.OrderId).Update("status", OrderConfirming)
					//}
				}
			case OrderConfirming:
			case OrderFinished:
			}
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, "更新成功")
	}

}

func (orderStep) Delete(c *gin.Context) {
	osid := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id=?", osid).Delete(&model.OrderStep{})
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "删除成功")
	}
}

func (orderStep) BatchDelete(c *gin.Context) {
	var params inout.BatchDeleteReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			tx.Where("id =?", id).Delete(&model.OrderStep{})
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "批量删除成功")
	}

}
