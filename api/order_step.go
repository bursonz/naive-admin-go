package api

import (
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
		if params.LockStatus != nil {
			err = orm.Update("lock_status", params.LockStatus).Error
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
		if params.Status != nil {
			err = orm.Update("status", params.Status).Error
			if err != nil {
				return err
			}
			// 获取当前步骤信息
			var current model.OrderStep
			orm.Select("order_id, task, status").Find(&current)
			var expectStatus, targetStatus int
			// 计算期望的状态和目标状态
			switch current.Task {
			case OrderStepTaskUnlock: // 执行
				expectStatus = OrderStepStatusExecuted
				targetStatus = OrderReviewing
			case OrderStepTaskUploadConfirmImage, OrderStepTaskRetrieveStatusValue: // 审核
				expectStatus = OrderStepStatusReviewed
				targetStatus = OrderConfirming
			case OrderStepTaskLock: // 确认
				expectStatus = OrderStepStatusConfirmed
				targetStatus = OrderFinished
			}
			// 如果当前步骤的状态等于期望的状态，且符合要求，则更新工单状态为目标状态
			if current.Status == expectStatus {
				var taskTotal, currentTotal int64
				t := tx.Model(&model.OrderStep{}).Where("order_id =?", current.OrderId)
				// 统计当前任务类型的所有步骤的数量
				t.Where("task =?", current.Task).Count(&taskTotal)
				// 统计当前任务类型的已完成步骤的数量
				t.Where("status =?", expectStatus).Count(&currentTotal)
				// 如果已完成步骤的数量等于总步骤数量，更新工单状态为审核中
				if currentTotal == taskTotal {
					err = tx.Model(&model.Order{}).Where("id =?", current.OrderId).Update("status", targetStatus).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, err)
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
		Resp.Succ(c, err)
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
