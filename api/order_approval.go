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
		orm.Where("order_id like ?", "%"+orderId+"%")
	}
	if approverId != "" {
		orm.Where("approver_id like ?", "%"+approverId+"%")
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
			Comment:    params.Comment,
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
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		var err error
		orm := tx.Model(&model.OrderApproval{}).Where("id=?", params.Id)
		if params.ApproverId != nil {
			err = orm.Update("approver_id", params.ApproverId).Error
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
		if params.Status != nil {
			err = orm.Update("status", params.Status).Error
			if err != nil {
				return err
			}
			// 更新工单状态
			var current model.OrderApproval
			orm.Select("order_id, status,sort").Find(&current)
			switch *params.Status {
			case OrderApprovalStatusApproved:
				var approvedTotal, total int64
				tx.Model(&model.OrderApproval{}).
					Where("order_id=?", current.OrderId).Count(&total).
					Where("status=?", OrderApprovalStatusApproved).Count(&approvedTotal)
				if total == approvedTotal {
					tx.Model(&model.Order{}).Where("id=?", current.OrderId).Update("status", OrderExecuting)
				}
			case OrderApprovalStatusRejected:
				tx.Model(&model.Order{}).Where("id=?", current.OrderId).Update("status", OrderRejected)
			}
		}
		return err
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, nil)
	}
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
func (orderApproval) BatchDelete(c *gin.Context) {
	var params inout.BatchDeleteReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			tx.Where("id =?", id).Delete(&model.OrderApproval{})
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "批量删除成功")
	}

}
