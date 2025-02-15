package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

type order struct{}

var Order = &order{}

func (order) Add(c *gin.Context) {
	var params inout.AddOrderReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		// 保存工单
		var newOrder = model.Order{
			DispatcherId: params.DispatcherId,
			StationId:    params.StationId,
			Status:       OrderApproving,
			OperatorId:   params.OperatorId,
		}
		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}
		// 保存工单审批
		for i, approval := range params.OrderApprovals {
			newApproval := model.OrderApproval{
				OrderId:    newOrder.ID,
				ApproverId: approval.ApproverId,
				Status:     OrderApprovalStatusApproving,
				Comment:    approval.Comment,
				Sort:       i + 1,
			}
			if err := tx.Create(&newApproval).Error; err != nil {
				return err
			}
		}
		// 保存工单步骤
		for i, step := range params.OrderSteps {
			newStep := model.OrderStep{
				OrderId:    newOrder.ID,
				Task:       step.Task,
				Sort:       i + 1,
				ReviewerId: step.ReviewerId,
				Status:     OrderStepStatusExecuting,
				LockId:     step.LockId,
				LockStatus: step.LockStatus,
				ImageUrl:   step.ImageUrl,
				Comment:    step.Comment,
			}
			if err := tx.Create(&newStep).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "")
	}

}

// Update 更新工单,只更新工单信息部分，不更新工单审批和工单步骤，工单审批和工单步骤的更新需要单独的接口
func (order) Update(c *gin.Context) {
	var params inout.PatchOrderReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(&model.Order{}).Where("id =?", params.Id)
	if params.DispatcherId != nil {
		orm = orm.Update("dispatcher_id", params.DispatcherId)
	}
	if params.OperatorId != nil {
		orm = orm.Update("operator_id", params.OperatorId)
	}
	if params.StationId != nil {
		orm = orm.Update("station_id", params.StationId)
	}
	if params.Status != nil {
		orm = orm.Update("status", params.Status)
	}
	if err := orm.Error; err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, "更新工单成功")
	}
}
func (order) List(c *gin.Context) {
	var data inout.OrderListRes
	// 查询条件 TODO 其他条件
	var id = c.DefaultQuery("id", "")
	var dispatcherId = c.DefaultQuery("dispatcherId", "")
	var stationId = c.DefaultQuery("stationId", "")
	var operatorId = c.DefaultQuery("operatorId", "")
	var approverId = c.DefaultQuery("approverId", "")
	var reviewerId = c.DefaultQuery("reviewerId", "")
	var status = c.DefaultQuery("status", "")
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var deleted = c.DefaultQuery("deleted", "")
	var all = c.DefaultQuery("all", "")
	// 条件查询
	var orm = db.Dao
	if deleted != "" {
		orm = orm.Unscoped()
	}
	orm = orm.Model(&model.Order{})
	if id != "" {
		orm = orm.Where("id like ?", id+"%")
	}
	if stationId != "" {
		orm = orm.Where("station_id like ?", stationId+"%")
	}
	if status != "" {
		orm = orm.Where("status like ?", status+"%")
	}
	if dispatcherId != "" {
		orm = orm.Where("dispatcher_id like ?", dispatcherId+"%")
	}
	if operatorId != "" {
		orm = orm.Where("operator_id like ?", operatorId+"%")
	}
	if approverId != "" {
		orm.Where("id in(?)", db.Dao.Model(&model.OrderApproval{}).Where("approver_id =?", approverId).Select("order_id"))
	}
	if reviewerId != "" {
		orm.Where("id in(?)", db.Dao.Model(&model.OrderStep{}).Where("reviewer_id =?", reviewerId).Select("order_id"))
	}
	// 查询总数
	orm.Count(&data.Total)
	// 分页查询
	if pageNo < 1 { // pageNo 小于1 时，查询所有
		orm.Find(&data.PageData)
	} else {
		orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)
	}
	if all != "" {
		// 遍历所有的工单，查询工单审批和工单步骤
		for i, o := range data.PageData {
			// 查询工单审批
			var approvals []model.OrderApproval
			db.Dao.Model(&model.OrderApproval{}).Where("order_id =?", o.ID).Find(&approvals)
			// 查询工单步骤
			var steps []model.OrderStep
			db.Dao.Model(&model.OrderStep{}).Where("order_id =?", o.ID).Find(&steps)
			// 将工单审批和工单步骤赋值给工单
			data.PageData[i].OrderApprovals = approvals
			data.PageData[i].OrderSteps = steps
		}
	}
	Resp.Succ(c, data)
}
func (order) Delete(c *gin.Context) {
	oid := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", oid).Delete(&model.Order{})
		tx.Where("order_id =?", oid).Delete(&model.OrderApproval{})
		tx.Where("order_id =?", oid).Delete(&model.OrderStep{})
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, err)
	}
}

func (order) BatchDelete(c *gin.Context) {
	var params inout.BatchDeleteReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			tx.Where("id =?", id).Delete(&model.Order{})
			tx.Where("order_id =?", id).Delete(&model.OrderApproval{})
			tx.Where("order_id =?", id).Delete(&model.OrderStep{})
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "批量删除成功")
	}

}
