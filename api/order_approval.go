package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

type orderApproval struct{}

var OrderApproval = &orderApproval{}

func (orderApproval) List(c *gin.Context) {
	var data inout.OrderApprovalListRes
	// 查询参数
	var orderId = c.DefaultQuery("orderId", "")
	var approverId = c.DefaultQuery("approverId", "")
	var status = c.DefaultQuery("status", "")
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var deleted = c.DefaultQuery("deleted", "")
	// 条件查询
	var orm = db.Dao
	if deleted != "" {
		orm = orm.Unscoped()
	}
	orm = orm.Model(&model.OrderApproval{})
	if orderId != "" {
		orm.Where("order_id=?", orderId)
	}
	if approverId != "" {
		orm.Where("approver_id=?", approverId)
	}
	if status != "" {
		orm.Where("status=?", status)
	}
	orm.Count(&data.Total) // 查询总数
	if pageNo < 1 {
		orm.Find(&data.PageData)
	} else {
		orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)
	}

}

func (orderApproval) Add(c *gin.Context) {
	var params inout.AddOrderApprovalReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		// 保存工单审批
		var newOrderApproval = model.OrderApproval{
			OrderId:    params.OrderId,
			ApproverId: params.ApproverId,
			Status:     params.Status,
			Sort:       params.Sort,
		}
		if err := tx.Create(&newOrderApproval).Error; err != nil {
			return err
		}
		// 更新工单状态为待审批，因为添加了一条审批记录
		if err := tx.Model(&model.Order{}).Where("id=?", params.OrderId).Update("status", 1).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, err)
	}
}

// Update 更新工单审批, 只能更新状态和审批人
func (orderApproval) Update(c *gin.Context) {
	var params inout.PatchOrderApprovalReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(&model.OrderApproval{}).Where("id=?", params.Id)
	if params.Status != nil {
		orm.Update("status", params.Status)
	}
	if params.ApproverId != nil {
		orm.Update("approver_id", params.ApproverId)
	}
	if params.Comment != nil {
		orm.Update("comment", params.Comment)
	}
	// TODO 检查工单审批状态，是否可以更新工单，进入执行状态
	// 待审批->已审批||已驳回
	Resp.Succ(c, err)
}

func (orderApproval) Delete(c *gin.Context) {
	oaid := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id=?", oaid).Delete(&model.OrderApproval{})
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, err)
	}
}
