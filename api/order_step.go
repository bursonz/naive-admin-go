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
		orm = orm.Where("order_id =?", orderId)
	}
	if task != "" {
		orm = orm.Where("task =?", task)
	}
	if operatorId != "" {
		orm = orm.Where("operator_id =?", operatorId)
	}
	if reviewerId != "" {
		orm = orm.Where("reviewer_id =?", reviewerId)
	}
	if status != "" {
		orm = orm.Where("status =?", status)
	}
	if lockId != "" {
		orm = orm.Where("lock_id =?", lockId)
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
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(&model.OrderStep{}).Where("id = ?", params.Id)
	if params.Task != nil {
		orm.Update("task", params.Task)
	}
	if params.Sort != nil {
		orm.Update("sort", params.Sort)
	}
	if params.ImageUrl != nil {
		orm.Update("image_url", params.ImageUrl)
	}
	if params.Comment != nil {
		orm.Update("comment", params.Comment)
	}
	if params.LockId != nil {
		orm.Update("lock_id", params.LockId)
	}
	if params.LockStatus != nil {
		orm.Update("lock_status", params.LockStatus)
	}
	if params.ReviewerId != nil {
		orm.Update("reviewer_id", params.ReviewerId)
	}
	if params.Status != nil {
		orm.Update("status", params.Status)
	}
	// TODO 检查工单步骤状态，是否可以更新工单，完成工单
	// 待执行||重新执行->待审核
	// 待审核->已审核||重新执行
	// 已审核->待关锁->已完成
	if params.Status != nil {
		orm.Update("status", params.Status)
	}
	Resp.Succ(c, err)

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
